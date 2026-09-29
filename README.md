# Warehouse Diagnostics CLI

A Go command-line tool for exploring operational warehouse data in SQL Server.

The project uses Microsoft's WideWorldImporters sample database as an unfamiliar transactional system for practicing schema discovery, SQL/T-SQL, Go database access, CLI tooling, and system design.

## Architecture

```text
terminal command
    ↓
cmd
    ↓
service
    ↓
repository
    ↓
database connection
    ↓
SQL Server / WideWorldImporters
```

- `cmd/` handles terminal arguments, validation, and output.
- `internal/diagnostics/` contains models, service functions, and diagnostic queries.
- `internal/database/` manages SQL Server connections.
- `main.go` starts the CLI and hands execution to the command layer.

## How Inventory Risk Is Defined

The word **risk** in this project has a specific meaning. It is a diagnostic rule used to identify order lines that may not have enough inventory available to finish picking.

The core calculation is:

```text
RemainingToPick = OrderedQuantity - PickedQuantity
```

An order line is part of the **picking backlog** when:

```text
PickedQuantity < OrderedQuantity
```

An order line becomes a **risky line** when it is still unfinished and the remaining quantity is greater than the current quantity on hand:

```text
PickedQuantity < OrderedQuantity
AND
RemainingToPick > QuantityOnHand
```

The metrics used throughout the CLI mean:

- **Ordered quantity** — the quantity requested on an order line.
- **Picked quantity** — the quantity that has already been picked for that order line.
- **Remaining to pick** — ordered quantity minus picked quantity.
- **Quantity on hand** — the inventory quantity currently stored in `Warehouse.StockItemHoldings`.
- **Picking backlog line** — an order line that has not been completely picked.
- **Risky line** — an unfinished order line whose remaining quantity is greater than the current quantity on hand.
- **Risky order** — a unique order that contains at least one risky line. An order is counted once even if several of its lines are risky.
- **Affected stock item** — a unique stock item that appears on at least one risky line.
- **Total remaining** — the sum of `RemainingToPick` across all risky lines included in the query.

For example:

```text
Ordered quantity:   120
Picked quantity:     20
Remaining to pick:  100
Quantity on hand:    40
```

Because 100 units remain to be picked but only 40 are currently on hand, that order line is classified as risky. The order containing it is counted as one risky order.

### Important data context

WideWorldImporters is a historical sample database. The project compares order-line data with the quantity currently present in `Warehouse.StockItemHoldings`. Because of that, a result labeled "risky" means **flagged by this project's diagnostic rule**, not necessarily a confirmed real-world stock shortage or failed allocation.

Commands such as `risk-by-year` and `recent-risk-summary` help separate historical records from the latest period represented in the dataset.

## Current Commands

### Picking backlog

```bash
go run main.go picking-backlog
```

Finds order lines that still have units left to pick and returns the order ID, date, stock item, remaining quantity to pick, and current quantity on hand.

### Inventory risk

```bash
go run main.go inventory-risk
```

Finds unfinished order lines where the remaining quantity to pick is greater than the current quantity on hand.

For this project:

```text
inventory risk = RemainingToPick > QuantityOnHand
```

### Order details

```bash
go run main.go order-details 21116
```

Investigates one order and returns each order line with:

- ordered quantity
- picked quantity
- remaining quantity to pick
- current quantity on hand

This makes it possible to move from a broad diagnostic such as `inventory-risk` into a specific order investigation.

### Risk summary

```bash
go run main.go risk-summary
```

Returns one aggregate summary of the dataset-wide inventory-risk results, including:

- unique risky orders
- risky order lines
- unique affected stock items
- total remaining units across risky lines
- oldest affected order date

This command uses SQL aggregates such as `COUNT`, `COUNT(DISTINCT ...)`, `SUM`, and `MIN` and returns a single summary row through `QueryRow()`.

### Items at risk

```bash
go run main.go items-at-risk
```

Groups inventory-risk results by stock item and returns:

- stock item name
- unique risky orders for that item
- risky order lines
- total remaining units
- current quantity on hand

Results are ordered by total remaining units so the largest item-level risks appear first.

This command introduces grouped aggregation with `GROUP BY`, making it possible to move from one overall risk summary into item-level analysis.

### Risk by year

```bash
go run main.go risk-by-year
```

Groups inventory-risk results by order year and returns:

- year
- unique risky orders
- risky order lines
- total remaining units

This command is used to distinguish historical risk from more recent operational risk and adds time-based grouping with `YEAR(...)` and `GROUP BY`.

### Recent risk summary

```bash
go run main.go recent-risk-summary
```

Summarizes inventory risk within the final 30 days of the dataset.

The command first finds the latest order date in WideWorldImporters, then uses that date as the dataset-relative endpoint for a 30-day window. It returns:

- window start
- window end
- unique risky orders
- risky order lines
- total remaining units

This avoids treating the sample database's historical records as if they were current-day production data.

The query introduces:

- a CTE with `WITH ... AS (...)`
- `MAX(...)` to find the latest order date
- `DATEADD(...)` to calculate the window start
- `CROSS JOIN` to make the one-row CTE result available to the main query
- `COALESCE(...)` to safely handle empty aggregate results

## Data Relationships

The current diagnostics use these WideWorldImporters tables:

```text
Sales.Orders
    ↓ OrderID
Sales.OrderLines
    ↓ StockItemID
Warehouse.StockItems
    ↓ StockItemID
Warehouse.StockItemHoldings
```

Key relationships:

```text
Sales.Orders.OrderID = Sales.OrderLines.OrderID

Sales.OrderLines.StockItemID = Warehouse.StockItems.StockItemID

Sales.OrderLines.StockItemID = Warehouse.StockItemHoldings.StockItemID
```

## Project Structure

```text
warehouse-diagnostics/
├── cmd/
│   ├── root.go
│   └── diagnostics.go
├── internal/
│   ├── database/
│   │   └── sqlserver.go
│   └── diagnostics/
│       ├── model.go
│       ├── service.go
│       └── repository.go
├── main.go
├── go.mod
└── go.sum
```

## Configuration

SQL Server runs locally in Docker.

The WideWorldImporters connection string is supplied through an environment variable rather than stored in source code:

```text
WWI_DB_URL
```

## Current Focus

The project is focused on the workflow used when entering an existing data-backed system:

```text
discover schema
    ↓
identify relationships
    ↓
write and validate SQL
    ↓
connect the query to Go
    ↓
expose it through the CLI
    ↓
diagnose an operational problem
```

The current diagnostic flow supports both detailed investigation and aggregate analysis:

```text
picking-backlog
    ↓
inventory-risk
    ↓
risk-summary
    ↓
items-at-risk / risk-by-year / recent-risk-summary
    ↓
order-details <OrderID>
```

Current SQL practice includes:

- multi-table joins
- parameterized queries
- calculated columns
- filtering with `WHERE`
- single-row aggregates with `QueryRow()`
- multi-row grouped results with `Query()`
- `COUNT(DISTINCT ...)`
- `SUM`, `MIN`, and `MAX`
- grouped aggregation with `GROUP BY`
- time-based grouping with `YEAR(...)`
- CTEs
- `DATEADD(...)`
- `CROSS JOIN`
- `COALESCE(...)`
- ordering aggregate results

Planned areas include:

- additional operational and historical diagnostics
- indexing and query performance
- execution plans
- automated tests
- ClickHouse for historical and analytical workloads
- AI-assisted diagnostic commands after the core SQL and ClickHouse workflows are established

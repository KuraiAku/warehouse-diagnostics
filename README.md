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

Planned areas include:

- summary and aggregate diagnostics
- more advanced SQL joins and aggregations
- indexing and query performance
- execution plans
- automated tests
- ClickHouse for historical and analytical workloads

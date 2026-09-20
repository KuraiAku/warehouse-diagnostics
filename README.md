# Warehouse Diagnostics CLI

A Go command-line tool for warehouse diagnostics and operational data exploration backed by SQL Server.

The project is built to practice tracing data through a real application flow:

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
SQL Server
```

The current stack includes:

- Go
- SQL / T-SQL
- SQL Server
- Linux / WSL terminal workflows
- Docker
- Database-backed CLI tooling
- Backend/system design
- Warehouse and transactional data

## Current Features

### View all inventory

```bash
go run main.go inventory-all
```

Returns warehouse, SKU, and quantity information from the local WarehouseDiagnostics database.

### Look up inventory by warehouse and SKU

```bash
go run main.go inventory ONT-01 BOX-1001
```

Filters inventory directly in SQL Server using warehouse code and SKU.

### Find low inventory

```bash
go run main.go inventory-low 50
```

Returns inventory records below the supplied quantity threshold using SQL-side filtering.

### Summarize orders by status

```bash
go run main.go orders-summary pending
```

Returns the supplied order status, order count, and total order value using an aggregate SQL query.

### Inspect picking backlog

```bash
go run main.go picking-backlog
```

Queries Microsoft's WideWorldImporters sample database to find order lines that still have units left to pick and returns:

- Order ID
- Order date
- Stock item name
- Remaining quantity to pick
- Current quantity on hand

This command joins operational data across:

```text
Sales.Orders
    ↓ OrderID
Sales.OrderLines
    ↓ StockItemID
Warehouse.StockItems
    ↓ StockItemID
Warehouse.StockItemHoldings
```

## Project Structure

```text
warehouse-diagnostics/
├── cmd/
│   ├── root.go
│   ├── inventory.go
│   ├── orders.go
│   └── diagnostics.go
├── db/
│   └── migrations/
├── internal/
│   ├── database/
│   │   └── sqlserver.go
│   ├── inventory/
│   │   ├── model.go
│   │   ├── service.go
│   │   └── repository.go
│   └── diagnostics/
│       ├── model.go
│       ├── service.go
│       └── repository.go
├── tools/
│   └── generate/
├── main.go
├── go.mod
└── go.sum
```

## Architecture

- `cmd/` handles CLI arguments and terminal output.
- `internal/inventory/` contains models, service functions, and SQL used by the original warehouse dataset.
- `internal/diagnostics/` contains diagnostics built against the larger WideWorldImporters sample database.
- `internal/database/` manages reusable SQL Server connections.
- `db/migrations/` contains the schema and seed data for the local WarehouseDiagnostics database.
- `tools/` contains development utilities rather than runtime CLI features.
- `main.go` starts the CLI and hands execution to the command layer.

The layers are kept separate so command handling, application logic, and database access can evolve independently.

## Databases

The CLI currently works with two SQL Server databases.

### WarehouseDiagnostics

A small local database created for learning the basic application flow.

Its schema includes:

- `warehouses`
- `products`
- `inventory`
- `orders`
- `order_items`
- `inventory_movements`

### WideWorldImporters

Microsoft's sample transactional database is used as a larger, unfamiliar system for practicing schema discovery, joins, operational diagnostics, and working with an existing database design.

The CLI currently uses tables including:

- `Sales.Orders`
- `Sales.OrderLines`
- `Warehouse.StockItems`
- `Warehouse.StockItemHoldings`

## Configuration

SQL Server runs locally in Docker.

Connection strings are supplied through environment variables rather than being stored in source code:

```text
WAREHOUSE_DB_URL
WWI_DB_URL
```

`WAREHOUSE_DB_URL` points to the local WarehouseDiagnostics database.

`WWI_DB_URL` points to WideWorldImporters.

The database package exposes reusable connection logic so different repositories can connect to the database they need without duplicating SQL Server setup code.

## Current Focus

The current focus is using a larger unfamiliar transactional database to practice the same workflow that would be needed when entering an existing system:

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
diagnose operational problems
```

Planned areas include:

- Additional operational diagnostic commands
- More advanced SQL joins and aggregations
- Indexing and query performance
- Execution plans
- Better automated testing
- ClickHouse for historical and analytical workloads

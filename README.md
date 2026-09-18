# Warehouse Diagnostics CLI

A Go command-line tool for warehouse diagnostics and operational data exploration backed by SQL Server.

The project focuses on building and tracing a small backend system using:

- Go
- SQL / T-SQL
- Linux terminal workflows
- SQL Server
- Database-backed CLI tooling
- Backend/system design
- Warehouse and transactional data

## Current Features

### View all inventory

```bash
go run main.go inventory-all
```

Returns warehouse, SKU, and quantity information from SQL Server.

### Look up inventory by warehouse and SKU

```bash
go run main.go inventory ONT-01 BOX-1001
```

The lookup is filtered directly in SQL Server using warehouse code and SKU rather than loading the full inventory table into Go.

### Find low inventory

```bash
go run main.go inventory-low 50
```

Returns inventory records below the supplied quantity threshold.

### Summarize orders by status

```bash
go run main.go orders-summary pending
```

Returns the supplied order status, order count, and total order value using an aggregate SQL query.

## Architecture

```text
CLI command
    ↓
Service
    ↓
Repository
    ↓
Database connection
    ↓
SQL Server
```

- `cmd/` handles CLI arguments and terminal output.
- `internal/inventory/` currently contains the data models, service functions, and repository queries used by the CLI.
- `internal/database/` manages the SQL Server connection.
- `db/migrations/` contains the schema and seed data.
- `main.go` starts the CLI and hands execution to the command layer.

The layers are kept separate so command handling, application logic, and database access can evolve independently as the project grows.

## Database

The current SQL Server schema includes:

- `warehouses`
- `products`
- `inventory`
- `orders`
- `order_items`
- `inventory_movements`

SQL Server runs locally in Docker with persistent storage.

Database credentials are supplied through the `WAREHOUSE_DB_URL` environment variable rather than being stored in source code.

## Current Focus

The current focus is moving beyond small seed data and learning how operational and analytical workloads should be handled as the system grows.

Planned areas include:

- Larger generated datasets
- More SQL-side filtering and aggregation
- Indexing and query performance
- Execution plans
- Additional warehouse diagnostic commands
- ClickHouse for historical and analytical workloads

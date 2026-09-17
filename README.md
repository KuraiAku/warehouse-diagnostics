# Warehouse Diagnostics CLI

A Go-based command-line tool for practicing warehouse system diagnostics using SQL Server.

The project is focused on:
- Go
- SQL / T-SQL
- Linux terminal workflows
- SQL Server
- Database-backed CLI tools
- Backend/system design
- Warehouse and transactional data

## Current Features

### View all inventory

go run main.go inventory-all

Returns warehouse, SKU, and quantity information from SQL Server.

### Look up inventory by warehouse and SKU

go run main.go inventory ONT-01 BOX-1001

The lookup is filtered directly in SQL Server rather than loading the entire inventory table into Go.

### Find low inventory

go run main.go inventory-low 50

Returns inventory records below the supplied quantity threshold.

## Architecture

CLI command
-> Service
-> Repository
-> Database connection
-> SQL Server

- cmd/ handles CLI input and output.
- internal/inventory/ contains inventory models, application logic, and database queries.
- internal/database/ manages the SQL Server connection.
- db/migrations/ contains the database schema and seed data.

## Database

The current SQL Server schema includes:
- warehouses
- products
- inventory
- orders
- order_items
- inventory_movements

SQL Server runs locally in Docker with persistent storage.

Database credentials are supplied through the WAREHOUSE_DB_URL environment variable rather than being stored in source code.

## Current Focus

I am currently extending the project with SQL-backed order diagnostics and aggregate queries.

Planned areas include:
- Order summaries and operational diagnostics
- Larger generated datasets
- Query/index performance
- Execution plans
- ClickHouse for analytical workloads
- Additional diagnostic commands

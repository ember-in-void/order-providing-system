# Frappuccino

A modern coffee shop management system built with Go and PostgreSQL. This application provides a RESTful API for managing menu items, inventory, orders, and sales aggregations.

## Features

- **Menu Management**: Create and manage coffee shop menu items
- **Inventory Tracking**: Monitor and update inventory levels for ingredients
- **Order Processing**: Handle customer orders and order items
- **Sales Analytics**: Generate aggregations and reports for business insights
- **Structured Logging**: Custom logger with file-based logging support
- **Clean Architecture**: Organized with separation of concerns (transport, service, storage layers)

## Tech Stack

- **Language**: Go 1.23
- **Database**: PostgreSQL 15
- **Container**: Docker & Docker Compose
- **Architecture**: Clean Architecture (Handler → UseCase → Repository)

## Project Structure

```
frappuccino/
├── internal/
│   ├── model/           # Data models and entities
│   ├── service/         # Business logic layer (usecases)
│   ├── storage/         # Database and repository layer
│   └── transport/       # HTTP handlers and routing
├── pkg/
│   └── logger/          # Custom logging utility
├── migrations/          # Database initialization scripts
├── logs/                # Application logs
└── main.go             # Application entry point
```

## Prerequisites

- Docker and Docker Compose installed on your system
- Port 8080 available for the application
- Port 5432 available for PostgreSQL (internal to Docker network)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/ember-in-void/frappuccino.git
cd frappuccino
```

### 2. Launch the Application

The easiest way to run the application is using Docker Compose:

```bash
docker-compose up --build
```

This command will:
- Build the Go application Docker image
- Start a PostgreSQL 15 database container
- Run database migrations (init.sql and insert.sql)
- Start the application on port 8080
- Mount the logs directory for persistent logging

### 3. Verify the Application

Once the containers are running, you should see:
```
Starting server on port 8080
```

The API will be available at `http://localhost:8080`

### 4. Stop the Application

To stop the application:

```bash
docker-compose down
```

To stop and remove volumes (including database data):

```bash
docker-compose down -v
```

## Running Without Docker

If you prefer to run the application locally without Docker:

### Prerequisites
- Go 1.23 or later
- PostgreSQL 15 or later

### Steps

1. **Start PostgreSQL** and create a database:
```sql
CREATE DATABASE frappuccino;
CREATE USER latte WITH PASSWORD 'latte';
GRANT ALL PRIVILEGES ON DATABASE frappuccino TO latte;
```

2. **Run migrations**:
```bash
psql -U latte -d frappuccino -f migrations/init.sql
psql -U latte -d frappuccino -f migrations/insert.sql
```

3. **Set environment variables**:
```bash
export DB_HOST=localhost
export DB_USER=latte
export DB_PASSWORD=latte
export DB_NAME=frappuccino
export DB_PORT=5432
```

4. **Install dependencies**:
```bash
go mod download
```

5. **Run the application**:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Configuration

The application uses environment variables for configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| DB_HOST | PostgreSQL host | db |
| DB_USER | Database user | latte |
| DB_PASSWORD | Database password | latte |
| DB_NAME | Database name | frappuccino |
| DB_PORT | Database port | 5432 |

## Development

### Project Dependencies

```bash
go mod tidy
```

### View Logs

Application logs are stored in the `./logs` directory and are automatically mounted when using Docker Compose.

## API Endpoints

The application exposes several API endpoints for managing:
- Menu items
- Inventory items
- Orders
- Aggregations/Analytics

Check the handler files in `internal/transport/handler/` for detailed endpoint information.

## Database Schema

Database schema and initial data are defined in:
- `migrations/init.sql` - Table definitions and schema
- `migrations/insert.sql` - Initial seed data


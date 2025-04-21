# HW-REST-API

Implementation of a simple API to manage applications, real-estate agents and customers.

## Dependencies

I've tried to use as few dependencies as possible. These are the ones I've chosen to go with:

- [echo](https://github.com/labstack/echo) as the web framework.
- [postgres](https://www.postgresql.org/) as the database.
- [validator.v10](https://github.com/go-playground/validator) to validate requests.
- [goose](https://github.com/pressly/goose) as the migration tool.

## Project structure

```
.
├── cmd/
│   └── main.go            # Main application entry point
├── config/                # Configuration
├── internal/
│   ├── app/               # Core application logic
│   ├── httpjson/          # Http handler adapters
│   ├── repository/        # Repository adapters
├── migrations/            # Database migration files and migration logic
```

## Running the application:

Create a `.env` file to set the needed env variables. Here's an example of a configuration that can be used:

```
SERVER_PORT=8080
SERVER_TIMEOUT_READ=3s
SERVER_TIMEOUT_WRITE=5s
SERVER_TIMEOUT_IDLE=5s

DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASS=topsecretpassword
DB_NAME=hw_db
```

The application can be run using docker by running the following command:

```
$ make build-run
```

## Running tests:

The test suite can executed by running the following command:

```
$ make run-test
```

## Endpoints:

| Method | Endpoint                                | Description                                                  |
| ------ | --------------------------------------- | ------------------------------------------------------------ |
| POST   | /agents                                 | Create a new agent                                           |
| PATCH  | /agents/{id}                            | Update a specific agent                                      |
| GET    | /customers/{customerId}/agent/{agentId} | Retrieve a specific agent connected with a specific customer |

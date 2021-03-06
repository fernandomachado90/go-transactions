# go-transactions
Sample API for financial transaction management written in Go on an implementation of hexagonal architecture.

## Commands

### `make setup`
Install dependency modules.

### `make format`
Format all files using `go fmt`.

### `make build`
Build source files into an executable binary named `bin`.

### `make test`                    
Execute all available tests.

### `make run`
Starts the server application locally on port `8080`.

### `make docker`
Starts the server application on a `Docker` image on port `8080`.

## Endpoints

#### `GET /healthcheck` 
Responds with status code `200` if the server is running.

#### `POST /accounts` 
Creates an account with the provided attributes.
###### request 
    {
        "document_number": "12345678900"
    }
###### response 
    {
        "account_id": "1",
        "document_number": "12345678900"
    }

#### `GET /accounts/:accountId` 
Searches for an account with the provided `:accountId`.
###### response 
    {
        "account_id": 1,
        "document_number": "12345678900"
    }

#### `POST /transactions` 
Creates a transaction with the provided attributes.
###### request
    {
        "account_id": 1,
        "operation_type_id": 4,
        "amount": 123.45
    }
###### response
    {
        "transaction_id": 1,
        "account_id": 1,
        "operation_type_id": 4,
        "amount": 123.45,
        "event_date": "2020-01-05T09:34:18.5893223"
    }

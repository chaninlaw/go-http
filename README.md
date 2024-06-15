# Simple HTTP server

## Quick Start
run following command:
```sh
go run cmd/main.go
```

if `Server starting on 8080` then you successfully to run the server.

## Setup Database
create an `.env` file
```sh
touch .env
echo -e 'DATABASE_URL=postgresql://username:password@localhost:5432/your_database_name\ADMIN_PASSWORD=12345678' > .env
```

## Setup Request Header
this repo have middleware for checking the request
add this `X-API-KEY: Hello world` to headers

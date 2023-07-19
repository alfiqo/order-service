## Getting started


Clone Project
```
$ git clone git@github.com:alfiqo/order-service.git

$ cd order-service && go mod tidy

$ cp .env.example .env
```
For running stack
```
$ docker compose up -d
```
Migrate database
```
$ migrate -database "mysql://user:pass@tcp(localhost:3306)/order_service" -path db/migrations up
```
Run App
```
go run main.go
```

Open API
```
specs
|-- customer
|   |-- customer.yml
|-- order
    |-- order.yml
```
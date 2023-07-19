## Getting started

```


git clone git@github.com:alfiqo/order-service.git

cd order-service && go mod tidy

migrate -database "mysql://root@tcp(localhost:3306)/order_service" -path db/migrations up

cp .env.example .env

go run main.go


// Open API
specs
|-- customer
|   |-- customer.yml
|-- order
    |-- order.yml
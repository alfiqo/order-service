package main

import (
	"order-service/config"
	"order-service/src/routes"
)

func main() {
	db := config.NewDB()
	defer config.CloseDB(db)

	routes.NewRoutes().Run()
}

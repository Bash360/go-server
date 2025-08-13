package main

import (
	"rest-api/app"
	o "rest-api/order"
	p "rest-api/product"

	_ "github.com/gorilla/mux"
)

func main() {
	app.InitializeRoutes(p.RegisterRoutes,
		o.RegisterRoutes)
	app.Run()
}

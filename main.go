package main

import (
	"rest-api/app"
	p "rest-api/product"

	_ "github.com/gorilla/mux"
)




func main(){
	 app.InitializeRoutes(p.RegisterRoutes)
   app.Run()
}

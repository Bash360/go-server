package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/app"
	"rest-api/product"

	_ "github.com/gorilla/mux"
)




func main(){
   product.RegisterRoutes(app.Server.Router)
	 http.Handle("/",app.Server.Router)
	 fmt.Println("Server listening on port "+app.Server.Port)
	 log.Fatal(http.ListenAndServe("localhost:"+app.Server.Port,nil))
 

}

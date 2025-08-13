package product

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/app"
	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)



func RegisterRoutes(router *mux.Router){
	 router.HandleFunc("/products",GetProducts).Methods("GET")
	 router.HandleFunc("/products/{id}",GetProduct).Methods("GET")
	 router.HandleFunc("/products",CreateProduct).Methods("POST")
}

func GetProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Product{}.GetProducts(app.Server.Connection))
}

func GetProduct(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)

	id,ok:=vars["id"]

	if !ok{
		log.Println("Id does not exist")
	}
	intId, err:=strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
	}
  product:=Product{}.GetProduct(intId,app.Server.Connection)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request){
	var product Product
	err :=json.NewDecoder(r.Body).Decode(&product)

	newProduct,err := Product{}.CreateProduct(&product, app.Server.Connection) 
	if  err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
  w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

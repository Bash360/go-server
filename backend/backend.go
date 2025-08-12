package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct{
	Id        int `json:"_"`
	Name      string `json:"name"`
	Inventory int    `json:"inventory"`
	Price     int    `json:"price"`
	ProductCode string `json:"productCode"`
	Status      string `json:"status"`   
}

type App struct{
	connection *sql.DB
	once sync.Once
	Port string
	Router *mux.Router
}
func(db *App) InitDB(){
	db.once.Do(func(){
		con, err := sql.Open("sqlite3","./practiceit.db")
	if err!=nil{
		log.Fatal(err.Error())
	}
  db.connection=con

	}) 
}

func(db *App)GetDBConnection() *sql.DB{
	return db.connection
}

func(app *App)GetProducts()[]Product{
	var products []Product
 rows, err := app.connection.Query("SELECT id, name, inventory, price, productCode,status FROM products")
 if err != nil {
	log.Fatal(err.Error())
 }
 defer rows.Close()

 for rows.Next(){
	var p Product

	rows.Scan(&p.Id, &p.Name, &p.Inventory, &p.Price,
		 &p.ProductCode,  &p.Status)
	products=append(products, Product{p.Id,
		p.Name,p.Inventory,p.Price,p.ProductCode, p.Status})
 }
 return products
}
 func (app *App) GetProduct(id int)Product{
	var p Product
   row:=app.connection.QueryRow(`SELECT id, name, inventory,
	  price, productCode, status FROM products where id=?`,id)
	err:= row.Scan(&p.Id,&p.Name,&p.Inventory,&p.Price, &p.ProductCode, &p.Status)

	if err != nil {
		log.Println("User doesn't exist",err)
	}
 return p
 }

 func (app *App)CreateProduct(p *Product)(Product,error){
	query:=`INSERT INTO products(productCode,name,inventory,price, status)
	 VALUES(?,?,?,?,?)`
	result, err:=app.connection.Exec(query,p.ProductCode,p.Name,p.Inventory,p.Price, p.Status)
	if err != nil{
		log.Println(err.Error())
	}
 id, err:= result.LastInsertId()
 if err !=nil{
	log.Println(err.Error())
	return Product{} ,err
 }
 p.Id=int(id)
	return *p, nil
 }
var  Server *App
func init(){
	Server=&App{Port: "3000"}
	Server.InitDB()
	Server.GetDBConnection()
  Server.Router = mux.NewRouter()
}

func Run(addr string){
	 Server.Router.HandleFunc("/products",GetProducts).Methods("GET")
	 Server.Router.HandleFunc("/products/{id}",GetProduct).Methods("GET")
	 Server.Router.HandleFunc("/products",CreateProduct).Methods("POST")
	 http.Handle("/",Server.Router)
	 fmt.Println("Server listening on port "+Server.Port)
	 log.Fatal(http.ListenAndServe(addr,nil))
}
func GetProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Server.GetProducts())
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
  product:=Server.GetProduct(intId)
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request){
	var product Product
	err :=json.NewDecoder(r.Body).Decode(&product)

	if err != nil{
		json.NewEncoder(w).Encode(err.Error())
	}

	newProduct,err := Server.CreateProduct(&product); 
	if  err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
	json.NewEncoder(w).Encode(newProduct)
}


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
	Id        int 
	Name      string 
	Inventory int
	Price     int
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
 rows, err := app.connection.Query("SELECT id, name, inventory, price FROM products")
 if err != nil {
	log.Fatal(err.Error())
 }
 defer rows.Close()

 for rows.Next(){
	var p Product

	rows.Scan(&p.Id, &p.Name, &p.Inventory, &p.Price)
	products=append(products, Product{p.Id,p.Name,p.Inventory,p.Price})
 }
 return products
}
 func (app *App) GetProduct(id int)Product{
	var p Product
   row:=app.connection.QueryRow(`SELECT id, name, inventory,
	  price FROM products where id=?`,id)
	err:= row.Scan(&p.Id,&p.Name,&p.Inventory,&p.Price)

	if err != nil {
		log.Println("User doesn't exist",err)
	}
 return p
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
	 Server.Router.HandleFunc("/products",HandlePost).Methods("POST")
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

func HandlePost(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode("This is a post end point haha")
}


package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

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
}

func Run(addr string){
	 http.HandleFunc("/products",GetProducts)
	 fmt.Println("Server listening on port "+Server.Port)
	 log.Fatal(http.ListenAndServe(addr,nil))
}
func GetProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Server.GetProducts())
}



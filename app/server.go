package app

import (
	"database/sql"
	"log"
	"sync"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

type App struct{
	Connection *sql.DB
	once sync.Once
	Port string
	Router *mux.Router
}

func(db *App) Init(){
	db.once.Do(func(){
		con, err := sql.Open("sqlite3","./practiceit.db")
	if err!=nil{
		log.Fatal(err.Error())
	}
	db.Router=mux.NewRouter()
  db.Connection=con

	}) 
}

func(db *App)GetDBConnection() *sql.DB{
	return db.Connection
}

var  Server *App

func init(){
	Server=&App{Port: "3000"}
	Server.Init()
 
 
}
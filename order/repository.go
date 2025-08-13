package order

import (
	"database/sql"
	"log"

_ "github.com/mattn/go-sqlite3"
)

type Order struct{
	ID int `json:"id"`
	CustomerName string  `json:"customerName"`
	Total int `json:"total"`
	Status string `json:"status"`
	Items []OrderItem `json:"items"`
}


type OrderItem struct{
	OrderID int `json:orderId`
	ProductID int `json:productId`
	Quantity int `json:quantity`
}

func getOrders()
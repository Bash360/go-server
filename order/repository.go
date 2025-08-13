package order

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Order struct {
	ID           int         `json:"id"`
	CustomerName string      `json:"customerName"`
	Total        int         `json:"total"`
	Status       string      `json:"status"`
	Items        []OrderItem `json:"items"`
}

type OrderItem struct {
	Order_id   int `json:"orderId"`
	Product_id int `json:"productId"`
	Quantity   int `json:"quantity"`
}

func (order Order) findAll(db *sql.DB) ([]Order, error) {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	orders := []Order{}
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.CustomerName, &o.Items, &o.Status, &o.Total); err != nil {
			log.Println(err.Error())
		}

		if err := o.findAllItems(db); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		orders = append(orders, o)

	}
	return orders, nil
}

func (o *Order) findAllItems(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM order_items WHERE order_id = ?", o.ID)
	if err != nil {
		log.Println(err.Error())
		return err

	}
	defer rows.Close()
	orderItems := []OrderItem{}

	for rows.Next() {
		var oi OrderItem
		if err := rows.Scan(&oi.Order_id, &oi.Product_id, &oi.Quantity); err != nil {
			log.Println(err.Error())
			return err
		}
		orderItems = append(orderItems, oi)
		o.Items = orderItems
	}
	return nil
}

func (o *Order) findOne(db *sql.DB) error {
	db.QueryRow(`SELECT customerName, total, 
	status FROM orders 
	where order_id =? `, o.ID).Scan(&o.ID, &o.Status, &o.CustomerName, &o.Total)

	err := o.findAllItems(db)
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) save(db *sql.DB) error {
	query := `INSERT INTO order(customerName, total, status) VALUES(?,?,?,?)`
	result, err := db.Exec(query, o.CustomerName, o.Total, o.Status)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	id, _ := result.LastInsertId()
	o.ID = int(id)
	return nil
}

func (oi *OrderItem) saveItem(db *sql.DB) error {
	query := `INSERT INTO order_items(order_id,product_id,quantity)
	VALUES(?,?,?)`
	_, err := db.Exec(query, oi.Order_id, oi.Product_id, oi.Quantity)
	if err != nil {
		log.Println(err.Error())
		return err

	}
	return nil
}

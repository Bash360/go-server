package product

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id          int    `json:"_"`
	Name        string `json:"name"`
	Inventory   int    `json:"inventory"`
	Price       int    `json:"price"`
	ProductCode string `json:"productCode"`
	Status      string `json:"status"`
}

func (pr Product) findAll(connection *sql.DB) []Product {
	var products []Product
	rows, err := connection.Query("SELECT id, name, inventory, price, productCode,status FROM products")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p Product

		rows.Scan(&p.Id, &p.Name, &p.Inventory, &p.Price,
			&p.ProductCode, &p.Status)
		products = append(products, Product{p.Id,
			p.Name, p.Inventory, p.Price, p.ProductCode, p.Status})
	}
	return products
}
func (pr *Product) findOne(connection *sql.DB) (Product, error) {
	var p Product
	row := connection.QueryRow(`SELECT id, name, inventory,
	  price, productCode, status FROM products where id=?`, pr.Id)
	err := row.Scan(&p.Id, &p.Name, &p.Inventory, &p.Price, &p.ProductCode, &p.Status)

	if err != nil {
		log.Println("User doesn't exist", err)
		return Product{}, err
	}
	return p, nil
}

func (p *Product) save(connection *sql.DB) (Product, error) {
	query := `INSERT INTO products(productCode,name,inventory,price, status)
	 VALUES(?,?,?,?,?)`
	result, err := connection.Exec(query, p.ProductCode, p.Name, p.Inventory, p.Price, p.Status)
	if err != nil {
		log.Println(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return Product{}, err
	}
	p.Id = int(id)
	return *p, nil
}

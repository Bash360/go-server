package product

import "rest-api/app"

var productRepository Product

func getAll() []Product {
	return productRepository.findAll(app.Server.Connection)
}

func getOne(id int) (Product, error) {
  pPointer:=&Product{Id:id}
	p, err := pPointer.findOne(app.Server.Connection)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func create(p *Product) (Product, error) {
	newProduct, err := p.save(app.Server.Connection)
	if err != nil {
		return Product{}, err
	}
	return newProduct, nil
}

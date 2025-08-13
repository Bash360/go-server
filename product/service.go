package product

import "rest-api/app"

var productRepository Product
func getProducts()[]Product{
return productRepository.GetProducts(app.Server.Connection)
}

func getProduct(id int)(Product, error){

	p,err := productRepository.GetProduct(id,app.Server.Connection)
	if err!=nil {
   return Product{}, err
	}
	return p, nil
}

func addProduct(p *Product)(Product,error){
	newProduct,err := productRepository.CreateProduct(p, app.Server.Connection)
	if err !=nil {
		return Product{}, err
	}
	return newProduct, nil
}
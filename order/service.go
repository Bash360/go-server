package order

import (
	"rest-api/app"
)

var orderRepo Order
var orderItem OrderItem

func getAll() ([]Order, error) {
	order, err := orderRepo.findAll(app.Server.Connection)
	return order, err
}

func getOne(id int) (Order, error) {
	order:=&Order{ID: id}
	err := order.findOne(app.Server.Connection)
	return orderRepo, err
}

func create(o *Order) error {
	err := o.save(app.Server.Connection)
	if err != nil {
		return err
	}
	return nil
}

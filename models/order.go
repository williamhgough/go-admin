package models

// Order struct general order information
type Order struct {
	id           int
	item         Product
	value        float64
	purchaseDate string
	customer     Customer
}

// CreateOrder - add new order record
func (o *Order) CreateOrder() {}

// GetOrder - add new order record
func (o *Order) GetOrder(id int) {}

// UpdateOrder - add new order record
func (o *Order) UpdateOrder(id int) {}

// DeleteOrder - add new order record
func (o *Order) DeleteOrder(id int) {}

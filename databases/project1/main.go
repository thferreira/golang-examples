package project1

import "github.com/google/uuid"

type Product struct {
	ID    string
	Name  string
	Price float64
}

func MyNewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {

}

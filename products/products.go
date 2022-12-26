package products

import "fmt"

type Product struct {
	id uint
	name string
	price float64
	description string
	category Category
}

type Category struct {
	name string
}

var products = []Product{}

func (p Product) save() {
	products = append(products, p)
	fmt.Println("Product added")
}

func (p Product) getAll() {
	for _, product := range products {
		fmt.Println(product)
	}
}

func (p Product) getById(id uint) Product {
	var findProduct Product
	for _, product := range products {
		if product.id == id {
			findProduct = product
		}
	}
	return findProduct
}

func Init()  {
	p := Product{
		id: 1,
		name: "a",
		price:   22.0,
		description: "a",
		category: Category{"NI"},
	}
	p.save()
	p.getAll()
	fmt.Println(p.getById(1))
}
package ecommerce

import "fmt"

const (
	small  = "SMALL"
	medium = "MEDIUM"
	big    = "BIG"
)

func newProduct(productType string, cost float64) Product {
	switch productType {
	case small:
		return Small{cost}
	case medium:
		return Medium{cost}
	case big:
		return Big{cost}
	}
	return nil
}

type Product interface {
	calculateMaintenance() float64
}

func (s Small) calculateMaintenance() float64 {
	return s.cost
}

func (m Medium) calculateMaintenance() float64 {
	return m.cost * 1.03
}

func (b Big) calculateMaintenance() float64 {
	return b.cost*1.06 + 2500
}

type Small struct {
	cost float64
}

type Medium struct {
	cost float64
}

type Big struct {
	cost float64
}

func details(p Product) {
	fmt.Println(p.calculateMaintenance())
}

func Init() {
	s := newProduct("SMALL", 200)
	details(s)
	m := newProduct("MEDIUM", 200)
	details(m)
	b := newProduct("BIG", 200)
	details(b)
}

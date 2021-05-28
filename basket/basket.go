package basket

func NewBasket() *Basket {
	return &Basket{
		products: make(map[string]float64),
	}
}

type Basket struct {
	products map[string]float64
}

func (b *Basket) AddItem(productName string, price float64) {
	b.products[productName] = price
}

func (b *Basket) GetBasketSize() int {
	return len(b.products)
}

func (b *Basket) GetBasketTotal() float64 {
	basketTotal := 0.0
	shippingPrice := 0.0

	// sum up product net prices
	for _,value := range b.products {
		basketTotal += value
	}

	// add VAT, 20%
	basketTotal = basketTotal * 1.2

	// calculate shipping price
	// Delivery for basket under €10 is €3
	// Delivery for basket over €10 is €2

	if basketTotal < 10 {
		shippingPrice = 3.0
	} else {
		shippingPrice = 2.0
	}

	return shippingPrice+basketTotal
}
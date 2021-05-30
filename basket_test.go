package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"pdfminion/basket"
)

// based upon: https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go

// shelf to store articles that are available for shopping
type shopping struct {
	shelf  *basket.Shelf
	basket *basket.Basket
}

func (sh *shopping) addProductToBasket(productName string) error {
	sh.basket.AddItem(productName, sh.shelf.GetProductPrice(productName))
	return nil
}

func (sh *shopping) iShouldHaveProductsInTheBasket(productCount int) error {
	if sh.basket.GetBasketSize() != productCount {
		return fmt.Errorf(
			"expected %d products but there are %d",
			sh.basket.GetBasketSize(),
			productCount,
		)
	}

	return nil
}

func (sh *shopping) theOverallBasketPriceShouldBe(basketTotal float64) error {
	if sh.basket.GetBasketTotal() != basketTotal {
		return fmt.Errorf(
			"expected basket total to be %.2f but it is %.2f",
			basketTotal,
			sh.basket.GetBasketTotal(),
			)
	}
	return nil
}

func (sh *shopping) addProductToShelf(product string, price float64) (err error) {
	sh.shelf.AddProduct(product, price)
	return nil
}


func InitializeBasketScenario(ctx *godog.ScenarioContext) {
	sh := &shopping{}

	// runs before a scenario is tested
	ctx.BeforeScenario(func(*godog.Scenario) {
		sh.shelf = basket.NewShelf()
		sh.basket = basket.NewBasket()
	})

	ctx.Step(`^I add the "([^"]*)" to the basket$`, sh.addProductToBasket)
	ctx.Step(`^I should have (\d+) products? in the basket$`, sh.iShouldHaveProductsInTheBasket)
	ctx.Step(`^the overall basket price should be €(\d+)$`, sh.theOverallBasketPriceShouldBe)
	ctx.Step(`^there is a "([^"]*)", which costs €(\d+)$`, sh.addProductToShelf)
}


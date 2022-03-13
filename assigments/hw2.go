package main

import (
	"fmt"
	"log"
)

func main() {
	world := struct {
		English string
		Spanish string
		French  string
	}{
		"world",
		"mundo",
		"monde",
	}
	fmt.Println("Hello, %s/%s/%s!", world.English, world.Spanish, world.Spanish)

	RegisterItem(Prices, "banana", 345)
	fmt.Printf("Prices: %v\n", Prices)

	cart := new(Cart)
	cart.AddItem("eggs")
	cart.AddItem("banana")
	fmt.Printf("cart.hasMilk() = %v\n", cart.hasMilk())
	fmt.Printf("cart.HasItem(%v) = %v\n", "bread", cart.HasItem("bread"))

	cart.AddItem("milk")
	cart.Checkout()
}

type Price int64

func (p Price) String() string {
	return fmt.Sprintf("$%.2f", float64(p)/100)
}

var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

func RegisterItem(prices map[string]Price, item string, price Price) {
	if _, ok := prices[item]; ok {
		log.Printf("item %v already has a price %v\n", item, price)
	}
	prices[item] = price
}

type Cart struct {
	Items      []string
	TotalPrice Price
}

func (c *Cart) hasMilk() bool {
	for _, item := range c.Items {
		if item == "milk" {
			return true
		}
	}
	return false
}

func (c *Cart) HasItem(item string) bool {
	for _, i := range c.Items {
		if i == item {
			return true
		}
	}
	return false
}

func (c *Cart) AddItem(item string) {
	if price, ok := Prices[item]; ok {
		c.Items = append(c.Items, item)
		c.TotalPrice += price
	} else {
		log.Printf("item %v has no price\n", item)
	}
}

func (c *Cart) Checkout() {
	fmt.Println("Cart list")
	for _, i := range c.Items {
		fmt.Println(i, Prices[i])
	}

	fmt.Printf("Total: %v\n", c.TotalPrice)

	c.Items = c.Items[:0]
	c.TotalPrice = 0
}

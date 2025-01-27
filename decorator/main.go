package main

import "fmt"

func main() {
	pizza := &VeggieMania{}

	// Add cheese to topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	// Add tomato to topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf("Price of veggieMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}

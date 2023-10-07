package main

import "fmt"


type Coffee interface { // Coffee represents the base coffee interface
	Price() float64
	Description() string
}


type ConcreteCoffee struct{} // ConcreteCoffee is the basic coffee without any toppings

func (c *ConcreteCoffee) Price() float64 {
	return 2.0
}

func (c *ConcreteCoffee) Description() string {
	return "Plain Coffee"
}


type MilkDecorator struct { // MilkDecorator is a decorator to add milk to coffee
	Coffee
}

func (m *MilkDecorator) Price() float64 {
	return m.Coffee.Price() + 1.0
}

func (m *MilkDecorator) Description() string {
	return m.Coffee.Description() + " with Milk"
}


type SugarDecorator struct { // SugarDecorator is a decorator to add sugar to coffi
	Coffee
}

func (s *SugarDecorator) Price() float64 {
	return s.Coffee.Price() + 0.5
}

func (s *SugarDecorator) Description() string {
	return s.Coffee.Description() + " with Sugar"
}

func main() {
	
	coffee := &ConcreteCoffee{} // Create a coffee without anything

	
	coffeeWithMilkAndSugar := &MilkDecorator{&SugarDecorator{coffee}} // Add milk and sugar decorators

	// Print the details and price
	fmt.Printf("Coffee: %s\n", coffeeWithMilkAndSugar.Description())
	fmt.Printf("Price: $%.2f\n", coffeeWithMilkAndSugar.Price())
}

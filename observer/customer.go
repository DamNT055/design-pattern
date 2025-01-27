package main

import "fmt"

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Println("Sending email to customer", c.id, "for item", itemName)
}

func (c *Customer) getID() string {
	return c.id
}

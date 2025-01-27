package main

func main() {
	shirtItem := newItem("Nike Shirts")

	observer1 := &Customer{id: "abc@gmail.com"}
	observer2 := &Customer{id: "def@gmail.com"}

	shirtItem.register(observer1)
	shirtItem.register(observer2)

	shirtItem.updateAvailability()
}

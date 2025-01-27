package main

import "fmt"

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) updateAvailability() {
	fmt.Println("Item is now in stock", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}
func (i *Item) deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}
func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}
func removeFromSlice(observerList []Observer, observerToMove Observer) []Observer {
	observerLength := len(observerList)
	for i, observer := range observerList {
		if observerToMove.getID() == observer.getID() {
			observerList[observerLength-1], observerList[i] = observerList[i], observerList[observerLength-1]
			return observerList[:observerLength-1]
		}
	}
	return observerList
}

package main

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

type EmailClient struct {
	id string
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending email - %s available from cient %s\n", value, eC.id)
}

func main() {
	nvidiaItem := NewItem("RTX 3800")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34dc",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailable()
}

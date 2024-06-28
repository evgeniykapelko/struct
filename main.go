package main

import "fmt"

type Animal struct {
	Age  int
	Type string
	Name string
}

type AnimalAction interface {
	getAge() int
	getType() string
}

func (a *Animal) getAge() int {
	return a.Age
}

func (a *Animal) getType() string {
	return a.Type
}

func (a Animal) voice() {
	fmt.Printf("I am a %s\n", a.Type)
}

func getName(a *Animal) {
	fmt.Printf("I am %s\n", a.Name)
}

func getCommonType(a AnimalAction) {
	fmt.Printf("I am a type %s\n", a.getType())
}

func main() {
	doge := &Animal{
		Age:  18,
		Type: "dog",
		Name: "Alan",
	}

	cat := &Animal{
		Age:  2,
		Type: "cat",
		Name: "Kiss",
	}

	getName(cat)
	getName(doge)

	getCommonType(doge)
	getCommonType(cat)

	doge.voice()
	cat.voice()
}

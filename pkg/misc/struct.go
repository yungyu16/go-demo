package main

import "fmt"

func main() {
	p1, p2 := NewPerson("songjialin")
	person := p1
	fmt.Println(p2)

	fmt.Println(person)
	getName := person.GetName()
	fmt.Println(person)
	fmt.Println(getName)
	fmt.Println(person.GetHeight())
	fmt.Println(person.GetWeight())
}

func NewPerson(name string) (*Person, Person) {
	person := Person{Name: name, Human: Human{Height: 180, Weight: 80}}
	person2 := Person{Human{Height: 1}, "送"}
	return &person, person2

}

type Person struct {
	Human
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

type Human struct {
	Height int
	Weight int
}

func (h Human) GetHeight() int {
	return h.Height
}

func (h Human) GetWeight() (w int, err error) {
	w = h.Weight
	return
}

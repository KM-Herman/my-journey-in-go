package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{ name string }

func (d Dog) Speak() string {
	return "Woof! for " + d.name
}

type Cat struct{ name string }

func (c Cat) Speak() string {
	return "Meow! for " + c.name
}

type Cow struct{ name string }

func (r Cow) Speak() string {
	return "Mhhaow! for " + r.name
}

func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main(){
	animals := []Speaker{
		Dog{name: "Shefferd"},
		Cat{name: "Kitty"},
		Cow{name: "Gaju"},
	}

	for _, animal := range animals{
		MakeSound(animal)
	}
}
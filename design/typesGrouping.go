package main

import "fmt"

type Cat struct {
	Name              string
	IsAnimal          bool
	CatSpecificFactor int
}

type Dog struct {
	Name              string
	IsAnimal          bool
	DogSpecificFactor int
}

type Speaker interface {
	speak()
}

func (d *Dog) speak() {
	fmt.Println("Woof! My name is a ", d.Name, ". I am an animal it is ", d.IsAnimal)
}

func (c *Cat) speak() {
	fmt.Println("Woof! My name is a ", c.Name, ". I am an animal it is ", c.IsAnimal)
}

func main() {
	speakers := []Speaker{
		&Dog{
			Name:              "Sharik",
			IsAnimal:          true,
			DogSpecificFactor: 4,
		},
		&Cat{
			Name:              "Murka",
			IsAnimal:          true,
			CatSpecificFactor: 4,
		},
	}

	for _, v := range speakers {
		v.speak()
	}
}

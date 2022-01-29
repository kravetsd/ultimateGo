package main

import "fmt"

type printer interface {
	print()
}

type user struct {
	name string
}

func (u user) print() {
	fmt.Printf("The user %s is printing\n", u.name)
}
func main() {
	us := user{"Dima"}
	us.print()

	entities := []printer{us, &us}

	us.name = "Dasha"

	for _, v := range entities {
		v.print()
	}
}

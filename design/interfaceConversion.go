package main

import "fmt"

type bike struct{}

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
}

func (bike) Move() {
	fmt.Println("Moving the bike.")
}

func (bike) Lock() {
	fmt.Println("Locking the bike.")
}

type MoveLocker interface {
	Mover
	Locker
}

var m Mover
var ml MoveLocker

func main() {

	ml = bike{}
	m = ml
	// ml = m This will not work here. Missing "Lock" method.

}

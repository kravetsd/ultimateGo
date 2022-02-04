package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Car struct{}

func (Car) String() string {
	str := "This is a Car"
	fmt.Println(str)
	return str
}

func (Cloud) String() string {
	str := "This is a Cloud"
	fmt.Println(str)
	return str

}

type Cloud struct{}

func main() {
	rand.Seed(time.Now().UnixNano())

	mvs := []fmt.Stringer{
		Car{},
		Cloud{},
	}

	for i := 0; i < 10; i++ {
		rn := rand.Intn(2)

		if v, is := mvs[rn].(Cloud); is {
			fmt.Println("Got Lucky:", v)
			continue
		}

		fmt.Println("Got Unlucky")
	}
}

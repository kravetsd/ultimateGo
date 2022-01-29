package main

import (
	"fmt"
)

type printer interface {
	print()
}

type canon struct {
	name string
}

func (c canon) print() {
	fmt.Printf("this is %s printing this message\n", c.name)
}

type hp struct {
	name string
}

func (h *hp) print() {
	fmt.Printf("this is %s printing this message\n", h.name)
}

func main() {
	cn := canon{"CAN-123-d"}
	hpp := hp{"HP-Laserjet-123"}
	ent := []printer{cn, &hpp}

	cn.name = "CAN-321-a"
	hpp.name = "LASERHJET-653-H"

	for _, v := range ent {
		v.print()
	}
}

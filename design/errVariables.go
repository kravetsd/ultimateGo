package main

import (
	"errors"
	"fmt"
)

var (
	ErrBadRequest = errors.New("Bad request!")
	ErrPageMoved  = errors.New("Page was mbed")
)

func main() {

	if err := Webcall(true); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad request happened")
			return

		case ErrPageMoved:
			fmt.Println("Page was moved")
			return

			// default:
			// 	fmt.Println(err)
			// 	return
		}
	}
	fmt.Println("Life is good")

}

func Webcall(b bool) error {
	if b {
		return ErrBadRequest
	}
	return ErrPageMoved
}

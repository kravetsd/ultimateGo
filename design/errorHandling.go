package main

import "fmt"

type error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func main() {
	if err := Webcall(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Life is good")
}

func Webcall() error {
	return New("Bad request")
}

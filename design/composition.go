package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Data struct {
	Line string
}

type Xenia struct {
	Host    string
	Timeout time.Duration
}

func (x *Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("Error reading data from Xenia")
	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

type Pilar struct {
	Host    string
	Timeout time.Duration
}

func (p *Pilar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

func Pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		if err := x.Pull(&data[i]); err != nil {
			return len(data[:i]), err
		}
	}
	return len(data), nil
}

func Store(p *Pilar, data []Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return len(data[:i]), err
		}
	}
	return len(data), nil
}

func main() {
	fmt.Println("This is a composition example.")
}

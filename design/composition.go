package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type System struct {
	Xenia
	Pillar
}

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

type Pillar struct {
	Host    string
	Timeout time.Duration
}

func (p *Pillar) Store(d *Data) error {
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

func Store(p *Pillar, data []Data) (int, error) {
	for i := range data {
		if err := p.Store(&data[i]); err != nil {
			return len(data[:i]), err
		}
	}
	return len(data), nil
}

func Copy(s *System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := Pull(&s.Xenia, data)
		if i > 0 {
			if _, err := Store(&s.Pillar, data[:i]); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
}

func main() {
	fmt.Println("This is a composition example.")

	sys := System{
		Xenia: Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Pillar: Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}

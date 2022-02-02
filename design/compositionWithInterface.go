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

type System struct {
	Puller
	Storer
}

type Storer interface {
	Store(d *Data) error
}

type Puller interface {
	Pull(*Data) error
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

func Pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return len(data[:i]), err
		}
	}
	return len(data), nil
}

func Store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return len(data[:i]), err
		}
	}
	return len(data), nil
}

func Copy(s *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := Pull(s, data)
		if i > 0 {
			if _, err := Store(s, data[:i]); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
}

func main() {
	sys := System{
		Puller: &Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Storer: &Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}
	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}

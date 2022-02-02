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

type Storer interface {
	Store(d *Data) error
}

type Puller interface {
	Pull(*Data) error
}

type PullStorer interface {
	Puller
	Storer
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

func Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := Pull(ps, data)
		if i > 0 {
			if _, err := Store(ps, data[:i]); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
}

//Now lets instanciate these interfaces:

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

type System1 struct {
	Pillar
	Xenia
}

// ----anothther example of storer and puller:

type Bob struct {
	Host    string
	Timeout time.Duration
}

func (b *Bob) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("Error reading data from Bob")
	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

type Ann struct {
	Host    string
	Timeout time.Duration
}

func (a *Ann) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

type System2 struct {
	Bob
	Ann
}

func main() {
	sys1 := System1{
		Xenia: Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Pillar: Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	sys2 := System2{
		Bob: Bob{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Ann: Ann{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys1, 3); err != io.EOF {
		fmt.Println(err)
	}

	if err := Copy(&sys2, 3); err != io.EOF {
		fmt.Println(err)
	}
}

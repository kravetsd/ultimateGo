package main

import (
	"fmt"
)

type file struct {
	name string
}

type pipe struct {
	name string
}

type reader interface {
	read(b []byte) (int, error)
}

func main() {
	fmt.Println("Hello interfaces")
	f := file{"myfile.txt"}
	p := pipe{"named_pipe"}

	retrieve(p)

	retrieve(f)

	bt := make([]byte, 100)
	cp(bt)
	fmt.Printf("%s\n", bt)

}
func cp(b []byte) {
	s := "this is string from cp method"
	copy(b, s)
}

func (f file) read(b []byte) (int, error) {
	s := "This is a file and some text in it"
	copy(b, s)
	return len(s), nil
}

func (p pipe) read(b []byte) (int, error) {
	r := "This is a pipe and some text in it"
	copy(b, r)
	return len(r), nil
}

func retrieve(r reader) error {
	sl := make([]byte, 4096)

	len, err := r.read(sl)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", sl[:len])
	return nil
}

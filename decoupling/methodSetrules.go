package main

import "fmt"

type duration int

func (d *duration) notify() {
	fmt.Printf("Sending notification in %d \n", d)
}

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Notify user %s via it's email %s\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}
func main() {
	us := user{"Dima", "admin@example.com"}
	sendNotification(&us)

	duration(42).notify()
}

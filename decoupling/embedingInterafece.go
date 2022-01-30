package main

import "fmt"

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Notify admin %s via it's email %s\n", a.name, a.email)
}

type notifier interface {
	notify()
}

func (u *user) notify() {
	fmt.Printf("Notify user %s via it's email %s\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	ad := admin{
		user{
			"John Smith",
			"jsmith@example.com",
		},
		"super",
	}

	sendNotification(&ad)
	// Notify user John Smith via it's email jsmith@example.com in case of inner is not overrided
	// Notify admin John Smith via it's email jsmith@example.com in case of an admin method overrides iuser's method
}

package main

import "fmt"

type RSVP struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*RSVP, 0, 10)

func main() {
	fmt.Println("Hello Asad")
}

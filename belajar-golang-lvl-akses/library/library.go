package library

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}

// Init function is called
var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "John Wick"
	Student.Grade = 3

	fmt.Println("--> library/library.go imported")
}

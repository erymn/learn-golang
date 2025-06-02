package main

import (
	. "belajar-golang-lvl-akses/libmodules"
	lm2 "belajar-golang-lvl-akses/libmodules2"
	"belajar-golang-lvl-akses/library"
	"fmt"
)

func main() {
	println("Original code tanpa prefix")
	library.SayHello("Eko")

	println("Call code from partial class")
	sayHelloFromPartial("Budi")

	println("Sample import module/package dengan prefix '. belajar-golang-lvl-akses/libmodules'")
	var s1 = Student{Name: "Eko", Grade: 2}
	fmt.Println("Nama :", s1.Name)
	fmt.Println("Grade :", s1.Grade)

	println("Sample import module/package dengan alias lm2 'lm2 belajar-golang-lvl-akses/libmodules2'")
	var p1 = lm2.Profile{Name: "Eko", Age: 30, Married: false}
	fmt.Println("Nama :", p1.Name)
	fmt.Println("Age :", p1.Age)
	fmt.Println("Married :", p1.Married)

	println("Fungsi init di library.go yang dipanggil di main.go")
	fmt.Printf("Name  : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)

}

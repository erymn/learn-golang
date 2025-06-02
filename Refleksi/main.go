package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 42
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}
	fmt.Println("Kalau cuma sebagai output tanpa proses bisa pakai interface{}")
	fmt.Println("nilai variabel :", reflectValue.Interface())

	fmt.Println("Casting dari interface{} ke int")
	var nilai = reflectValue.Interface().(int)
	fmt.Println("nilai variabel setelah casting int [var nilai = reflectValue.Interface().(int)]:", nilai)

	fmt.Println("Lakukan uji coba method getPropertyInfo()")
	fmt.Println("dari struct student")
	var s = student{Name: "wick", Grade: 2}
	s.getPropertyInfo()

	fmt.Println("Set Property lewat code")
	var s1 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue1 = reflect.ValueOf(s1)
	fmt.Println(reflectValue1)
	var method = reflectValue1.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})
	fmt.Println("nama :", s1.Name)

}

package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama\t\t:", reflectType.Field(i).Name)
		fmt.Println("tipe data\t:", reflectType.Field(i).Type)
		fmt.Println("nilai\t\t:", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

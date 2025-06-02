package main

import "fmt"

var person = []map[string]interface{}{
	{"name": "Wick", "age": 23},
	{"name": "Ethan", "age": 23},
	{"name": "Bourne", "age": 22},
}

var fruits = []interface{}{
	map[string]interface{}{"name": "strawberry", "total": 10},
	[]string{"manggo", "pineapple", "papaya"},
	"orange",
}

func main() {
	fmt.Println("Kombinasi Slice, Map, dan Interface{}")
	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	fmt.Println("Kombinasi Map, dan Interface{}")
	for _, each := range fruits {
		fmt.Println(each)
	}

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}
	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	var data map[string]interface{}
	data = map[string]interface{}{
		"name":    "ethan hunt",
		"grade":   2,
		"message": "awesome",
		"breakfast": []string{
			"apple",
			"manggo",
			"banana",
		},
	}

	fmt.Println(data)

	fmt.Println("Alternatif interface")
	var dataAlias map[string]any

	dataAlias = map[string]any{
		"name":      "ethan hunt any",
		"grade":     23,
		"breakfast": []string{"apple1", "manggo", "banana"},
	}

	fmt.Println(dataAlias)

	fmt.Println("Casting data tipe interface{} atau any")
	var secret2 interface{} = "ethan hunt"
	var name string = secret2.(string)
	fmt.Println(name)

	secret2 = 100
	var value int = secret2.(int) * 100
	fmt.Println("Nilai hasil kali 100 adalah: ", value)

	secret2 = []string{"apple", "manggo", "banana"}
	var fruits = strings.Join(secret2.([]string), ", ") // join array string dengan koma jad
	fmt.Println(fruits)

	fmt.Println("Casting Variabel Interface Kosong Ke Objek Pointer")
	var secret3 interface{} = &person{name: "wick", age: 27}
	var name3 = secret3.(*person).name
	fmt.Println(name3)

}

type person struct {
	name string
	age  int
}

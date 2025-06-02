package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	fmt.Println("Memakai fungsi goroutine")
	go print(50, "halo")

	fmt.Println("Pemakaian biasa tanpa goroutine")
	print(15, "apa kabar")

	var input string
	fmt.Scanln(&input)
}

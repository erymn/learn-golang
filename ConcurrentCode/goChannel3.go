package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMessage(messages)
	}

}

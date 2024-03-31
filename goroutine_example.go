package main

import (
	"fmt"
	"time"
)

func HelloWorld(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	go HelloWorld(19, "hallo")
	HelloWorld(5, "baik kah ?")

	time.Sleep(2 * time.Second)

	// var input string
	// fmt.Scanln(&input)
}
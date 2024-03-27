package main

import (
	"fmt"
	"time"
)

func HelloWorld(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
	time.Sleep(5 * time.Second)
}

func main() {
	go HelloWorld(25, "hallo")
	HelloWorld(10, "baik kah ?")

	var input string
	fmt.Scanln(&input)
}
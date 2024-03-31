package main

import "fmt"

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	
	var messsages = make(chan string)

	for _,each := range []string{"baihaqi","patrict","putra"}{
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messsages <- data
		}(each)
	}

	for i := 0; i < 3; i++{
		printMessage(messsages)
	}
}
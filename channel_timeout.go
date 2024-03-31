package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendData(ch chan <- int) {
	randomize := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++{
		ch <- i
		fmt.Println(time.Duration(randomize.Int()%10+1)* time.Second)
		time.Sleep(time.Duration(randomize.Int()%10+1)* time.Second)
	}
}

func retreiveData(ch <- chan int){
	loop:
	for {
		select{
		case data := <- ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <- time.After(time.Second * 5):
			fmt.Println("timeout, ga aktivitas gan !")
			break loop
		}
	}
}

func main() {

	// runtime.GOMAXPROCS(2)
	var message = make(chan int)

	go sendData(message)
	retreiveData(message)
}
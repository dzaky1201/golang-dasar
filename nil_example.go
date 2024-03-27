package main

import "fmt"

func employeIdentity(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	//nill hanya bisa digunakan di :

	/*
		- interface
		- function
		- map
		- slice
		- pointer
		- channel
	*/

	data1 := employeIdentity("")

	if data1 == nil {
		fmt.Println("Datanya kosong gan !")
	} else {
		fmt.Println(data1)
	}

}
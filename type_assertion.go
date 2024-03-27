package main

import "fmt"


func random()interface{}{
	return "Ok"
}

func main(){
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // akan panic
	// fmt.Println(resultInt)

	switch value := result.(type){

	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown")
	}
}
	
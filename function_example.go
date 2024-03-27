package main

import (
	"fmt"
	"strings"
)

func helloWorld(fisrtName string,lastName string){
	fmt.Println("Hello World", fisrtName, lastName)
}

func convertToLowerCase(result string) string{
	return strings.ToLower(result)
}

func getFullname()(fisrtname string, lastname string){
	return "Joni", "Juliansyah"
}


// variadic function
// variadic digunakan ketika kita memiliki sebuah param dengan tipe data yang sama
func listImage(images ...string) string{
	for _, data := range images{
		println(data)
	}

	return "sucess"
}

func sayHello(person string) string{
	return "Hei bro" + person
}

func helperToSum(value1 int, value2 int)int{
	return value1 + value2
}

// function as param
func thisIsFunctionParam(sample string, convertToLowerCase func(string) string){
	fmt.Println("this is convert to lower case : ", convertToLowerCase(sample))
}

type Sum func(int, int) int

func sumNumberResult(value1 int, value2 int, sum Sum){
	sumNumberValue := sum(value1,value2)
	fmt.Printf("Ini adalah hasil penjumlahan : %d \n", sumNumberValue)
}

func incrementNumberResult(value1 int, value2 int, sum Sum){
	sumNumberValue := sum(value1,value2)
	fmt.Printf("Ini adalah hasil pengurangan : %d \n", sumNumberValue)
}

func messageSuccess(){
	fmt.Println("Selesai melakukan looping")
}

func closedApp(){
	messageSuccess()
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error :", message)
	}else{
		fmt.Println("Error : 0")
	}
}

func loopingEmployee(error bool){

	//implement defer function
	// defer messageSuccess()

	defer closedApp()

	if error {
		panic("Error gan karena errornya true")
	}


	employes := []string{
		"Andri",
		"Jhone",
		"Due",
		"Pian",
	}

	for _, data := range employes {
		fmt.Println(data)
	}
}

func main() {
	helloWorld("Dzaky", "Haidar")
	result := convertToLowerCase("DZAKY")
	fisrtName, _ := getFullname()

	// slice for param listImage func
	listImage2 := []string{"image1.png","image2.png", "image3.png"}

	list := listImage("gambar.png", "gambar2.png")
	list2 := listImage(listImage2...)

	sayHi := sayHello
	fmt.Println(sayHi("Udin"))

	converter := convertToLowerCase
	thisIsFunctionParam("POLICE", converter)

	thisIsFunctionParam("JAIL", convertToLowerCase)

	sumNumberResult(10,10, helperToSum)

	// penggunaan anonymous function cara ke 1
	pengurangan := func (value1 int, value2 int) int {
		return value1 - value2
	}
	incrementNumberResult(90,10,pengurangan)

	// penggunaan anonymous function cara ke 2
	incrementNumberResult(85,11,func(value1 int, value2 int)int{
		return value1-value2
	})


	fmt.Println(result)
	fmt.Println(fisrtName)
	fmt.Println(list)
	fmt.Println(list2)
	loopingEmployee(false)

	//nill hanya bisa digunakan di :

	/* 
	- interface 
	- function
	- map
	- slice
	- pointer
	- channel
	*/
	
}
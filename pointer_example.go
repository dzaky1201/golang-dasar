package main

import "fmt"

func pointerUse(original *int, value int){
	*original = value // *original artinya kita mau ambil valuenya dari param tersebut bukan nilai pointernya
}

type Man struct{
	Name string
	Age int
}

func (man *Man) GetName()string{
	return "Mr." + man.Name
}

func main(){
	var a = "hello" // <--- ini adalah variable biasa
	var b *string = &a  //<--- variable pointer -> harus diisi dengan sebuah pointer juga
	var c = a

	fmt.Println("value variable a :", a)
	fmt.Println("memori variable a :", &a)

	fmt.Println("value variable b :", *b)
	fmt.Println("memori variable a :", b)
	fmt.Println("")

	a = "hello revision"

	fmt.Println("value variable a :", a)
	fmt.Println("memori variable a :", &a)

	fmt.Println("value variable b :", *b)
	fmt.Println("memori variable a :", b)

	fmt.Println("value variable c :", c)
	fmt.Println("memori variable c :", &c)

	var numbers = 10
	fmt.Println("")
	fmt.Println("nilai sebelum :", numbers)

	pointerUse(&numbers, 12)

	fmt.Println("nilai after :", numbers)

	dzaky := Man{"Dzaky"}
	fmt.Println(dzaky.GetName())
	
}
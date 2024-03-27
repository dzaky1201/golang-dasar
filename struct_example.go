package main

import (
	"fmt"
	"strconv"
)


type Identity interface {
	GetFullIdentity()string
}

// ini penulisan struct
type Animal struct{
	Name, Gender string
	Age int
	Adult bool
}

// method struct animal
func (animal Animal) sayHiToPet(){
	fmt.Println("hi, welcome to the jungle", animal.Name)
}

func (animal Animal) checkAgePet(){
	fmt.Printf("Umurnya adalah : %d\n", animal.Age)
}

func (animal Animal) GetFullIdentity() string{
	age := strconv.Itoa(animal.Age)
	return animal.Name + age + animal.Gender
}

func main()  {
	var pet Animal
	pet.Name = "Frank"
	pet.Age = 12
	pet.Adult = true

	// struct literals option 1
	pet2 := Animal{
		Name: "ciks",
		Gender: "Female",
		Age: 8,
		Adult: false,
	}

	// struct literals option 2
	pet3 := Animal{"bro","Male",70,true}

	pet4 := pet2

	pet4.Age = 19



	fmt.Println(pet)
	fmt.Println(pet2)
	fmt.Println(pet3)
	fmt.Println("pembatas gan")
	fmt.Println(pet2)
	fmt.Println(pet4)


	pet3.sayHiToPet()
	pet3.checkAgePet()
	fmt.Println(pet3.GetFullIdentity())

	// buat 2 buah struct persegi dan persegi panjang
	// buat 1 interface bangunruang dengan method luas dan keliling
	// 2 buah struct yang sudah dibuat implement interface bangunruang



}
package main

import "fmt"

func main(){
	nilai := 80
	nilai = 70
	if nilai > 70 {
		fmt.Println("Lulus")
	}else if nilai == 70 {
		fmt.Println("Lulus Bersyarat")
	}else{
		fmt.Println("Mengulang")
	}

	fullname := "Michael"

	
	if length := len(fullname); length > 5 {
		fmt.Println("To many char")
	}else if length <= 5 {
		fmt.Println("Normal name")
	}

	switch fullname {
	case "Saif":
		fmt.Println("Halo saif, selamat datang")
	case "Husen":
		fmt.Println("Halo udin")
	default:
		fmt.Println("Maneh teh saha ?")
	}

	switch length := len(fullname); length > 5{
	case true:
		fmt.Println("To many char")
	case false:
		fmt.Println("Normal name")
	}

	switch {
	case nilai > 70:
		fmt.Println("Lulus")
	case nilai == 70:
		fmt.Println("Lulus bersyarat")
	}

}
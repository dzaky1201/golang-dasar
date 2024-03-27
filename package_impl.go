package main

import "go-dasar/subfolder1"

func main() {
	car1 := subfolder1.Car{Merk: "Toyota", Type: "Veloz", Year: "2020"}
	car1.Detail()

	subfolder1.Create()
}
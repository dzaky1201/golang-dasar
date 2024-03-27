package main

import "fmt"

func main(){
	// var names10 = [8]int{
	// 	90,
	// 	20,
	// 	89,
	// 	78,
	// 	19,
	// 	80,
	// 	77,
	// 	96,
	// }
	
	// tipe data slice
	/* ketika mau slice array ditengah-tengah 
		[indexAwal:indexAkhir + 1]

		ketika mau slice array dari tengah - index terakhir
		[indexAwal:]

		ketika mau slice array dari awal - index spesifik
		[:indexakhir + 1]

		membuat slice baru kosong
		make([]tipedata, length, capacity)

		length = panjang dari slice
		capacity = kapasitas menambahkan data pada slice menggunakan append
	*/
	// slice1 := names10[4:6]
	// slice2 := names10[2:]
	// slice3 := names10[:4]

	// slice4 := names10[:]
	// slice5 := append(slice4, 97, 98, 99)
	slice6 := make([]string, 3, 4)
	slice6[0] = "Asep"
	slice6[1] = "Cecep"
	slice6[2] = "ipul"

	slice7 := append(slice6, "bla", "blu")




	// fmt.Print("Ini contoh slice 1: ", slice1, "\n")
	// fmt.Print("Ini contoh slice 2: ", slice2, "\n")
	// fmt.Print("Ini contoh slice 3: ", slice3, "\n")
	// fmt.Print("Ini contoh slice 4: ", slice4, "\n")
	// fmt.Print("Ini contoh slice 5: ", slice5, "\n")
	fmt.Printf("Ini contoh slice 6: %s \n", slice6)
	fmt.Printf("Ini contoh slice 7: %s \n", slice7)
	fmt.Printf("Capacity 6: %d \n", cap(slice6))
	fmt.Printf("Capacity 7: %d \n", cap(slice7))

	fmt.Println(slice6)
	fmt.Println(slice7)

	array1 := [...]int{
		1,
		2,
		5,
	}

	fmt.Println(array1)

	// array2 := [2]int{

	// }

	// sliceExam := []int{

	// }

}
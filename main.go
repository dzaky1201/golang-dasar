package main

import "fmt"

func main() {
	
	const numbers = 190

	const (
		firstname = "Saipul"
		lastname = "Haidar"
	)

	// operator matematika
	x := 10
	y := -90

	result1 := x * y
	result2 := x + y
	result3 := x - y
	result4 := x / y

	fmt.Println(numbers)
	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	var a = 10
	var b = 67


	/*
		ini contoh penjumlahan
	*/

	//before (tanpa augmented assigment)
	a = a + 10
	// after (menggunakan augmented assgiment)
	a += 10
	fmt.Print("Ini hasil penjumlahan ",a,"\n")

	/*
		ini contoh perkalian
	*/
	b *= 67
	b *= a
	fmt.Print("Ini hasil perkalian ", b, "\n")

	// unary operator
	/*
		terdiri dari => ++ (Increment) , -- (decrement) , -(Negative) , + (Positive), ! (Negasi => nilai kebalikan [Boolean])
	*/
	c := 2
	d := 7

	// c = c + 1
	c++

	// d = d - 1
	d--

	fmt.Print("Ini hasil increment ", c, "\n")
	fmt.Print("Ini hasil decrement ", d, "\n")

	school1 := "Binus"
	school2 := "uph"

	// contoh penggunaan unary operator (!)
	result := school1 != school2
	fmt.Println(result)

	// hasil dari unary operator
	fmt.Println(!result)

	/* 
		Operator perbandingan
		1. > artinya lebih dari
		2. < artinya kurang dari
		3. <= artinya kurang dari sama dengan
		4. >= artinya lebih dari sama dengan
		5. == sama dengan
		6. != tidak sama dengan
		hasil dari semua perbandingan adalah boolean (true dan fales)
	*/

	number3 := 90
	number4 := 87
	result5 := number3 > number4

	number5 := 87
	result6 := number5 >= number4
	
	fmt.Println(result5)
	fmt.Println(result6)

	/*
		Operasi Boolean => AND,OR,NEGATION
		operator AND => &&
		operator OR => ||
		operator NEGATION => !

		tabel penggunaan AND (&&)
		nilai1    operator   nilai2   result
		true	   &&		  true		true
		true       &&		  False     False
		False	   &&         true		False
		False	   &&		  False     False


		tabel penggunaan OR (||)
		nilai1    operator   nilai2   result
		true	   ||		  true		true
		true       ||		  False     True
		False	   ||         true		True
		False	   ||		  False     False

		tabel penggunaan OR (!)
	    operator   nilai2   result
		!			true	false
		!			false	true
		*/

		var names10 = [4]int{
			90,
			20,
			89,
			78,
		}

		fmt.Println(names10)
		fmt.Println(len(names10))
		fmt.Println(names10[2])

		names10[1] = 1

		fmt.Println(names10)

		// tipe data slice
		slice1 := names10[2:4]
		fmt.Print("Ini contoh slice 1: ", slice1, "\n")



	

}
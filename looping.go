package main

import "fmt"

func main() {
	// studentName := [...]string{
	// 	"Andi",
	// 	"Joko",
	// 	"Toni",
	// 	"Sultan",
	// }

	// for i := 0; i < len(studentName); i++ {
	// 	fmt.Println(studentName[i])
	// }

	// for _, name := range studentName {
	// 	fmt.Println(name)
	// }

	// for i := 0; i <= 20 ; i++{
	// 	if i % 2 != 0{
	// 		continue
	// 	}

	// 	if i == 11{
	// 		break
	// 	}

	// 	fmt.Println(i)
	// }

	for i := 0; i < 5; i++{
		for x := i; x < 5; x++{
			fmt.Print(x, " ")
		}
		fmt.Println()
	}
}
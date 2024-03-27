package main

import "fmt"

func main(){

	identity := map[string]string{
		"fullname": "ahmad baihaqi",
		"address": "Jakarta",
	}


	fmt.Println(identity)
	fmt.Println(identity["fullname"])
	fmt.Println(identity["address"])
	fmt.Println(len(identity))

	identity["address"] = "Bogor"

	fmt.Println(identity["address"])

	identity2 := make(map[string]string)
	identity2["family"] = "Baihaqi Fams"
	identity2["hobbies"] = "Parkour"
	identity2["hobbies"] = "Swimming"
	fmt.Printf("Ini contoh map dengan 2 value yang sama : %s\n", identity2)


}
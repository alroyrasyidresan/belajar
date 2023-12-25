package main

import "fmt"

func main()  {
	type NoKTP string

	var KTPAlroy NoKTP = "1234567890123456"
	var contoh string = "2222222222222222"

	var contohKTP = NoKTP(contoh)

	fmt.Println(KTPAlroy)	
	fmt.Println(contohKTP)

}
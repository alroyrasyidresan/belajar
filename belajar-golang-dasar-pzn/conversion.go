package main

import "fmt"

func main()  {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai64)
	fmt.Println(nilai16) //akan balik kembali ke nilai terbawah dari int16 lalu naik sebanyak kelebihannya


	var name = "Alroy Rasyid Resan"
	var e = name[0] //e adalah uint8 atau byte
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
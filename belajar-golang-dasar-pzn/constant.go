package main

import "fmt"

func main() {
	//kalau tidak dipakai tidak masalah
	//tidak dapat diubah isinya
	const firstName string = "Alroy"
	const lastName = "Rasyid Resan"

	//coba ubah isinya
	// firstName = "Candra"

	fmt.Println(firstName)
	fmt.Println(lastName)

	//deklarasi multiple constant
	const (
		constant1 = "value1"
		constant2 = 123
		constant3 = true
	)

	fmt.Println(constant1)
	fmt.Println(constant2)
	fmt.Println(constant3)
}

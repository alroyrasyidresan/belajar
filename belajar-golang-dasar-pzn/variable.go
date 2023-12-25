package main

import "fmt"

func main() {
	//semua variable harus digunakan
	var nama string

	nama = "Alroy Rasyid Resan"
	fmt.Println(nama)

	nama = "Alroy"
	fmt.Println(nama)

	//kalau langsung diinisialisasi maka tidak perlu tipe data
	var umur = 21
	var alamat string = "Jl. Pahlawan No. 1" //tapi tetap bisa

	fmt.Println(umur)
	fmt.Println(alamat)

	//untuk tanpa deklarasi tanpa var
	nama2 := "Alroy Rasyid Resan"
	fmt.Println(nama2)

	nama2 = "Alroy"
	fmt.Println(nama2)

	//deklarasi multiple variable
	var (
		firstName  = "Alroy"
		middleName = "Rasyid"
		lastName   = "Resan"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(firstName, middleName, lastName)
}

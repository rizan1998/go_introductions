package vargo

import (
	"fmt"
)

func Vartipe() {
	// tidak boleh deklarasi variable dengan sembarang ukuran tipe data karena berpeganruh pada memory
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// tipe data tanda ukuran otomatis akan di tentuknan oleh compiler
	var decimalNUmber = 3.14
	fmt.Printf("bilangan desimal: %f\n", decimalNUmber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNUmber)


	// tipe data boolean
	var exist bool = true
	fmt.Printf("apakah ada: %t\n", exist)

	// tipe data string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var message2 = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message2)

	// nil
	/*
	Zero value dari string adalah "" (string kosong).
	Zero value dari bool adalah false.
	Zero value dari tipe numerik non-desimal adalah 0.
	Zero value dari tipe numerik desimal adalah 0.

	pointer
	tipe data fungsi
	slice
	map
	channel
	interface kosong atau any (yang merupakan alias dari interface{})
	*/

	// konstanta tipe data yang tidak dapat diubah
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	// multi constanta
	const(
		square          = "kotak"
		isToday bool    = true
		numeric uint8   = 1
		floatNum        = 2.2
	)

	// Ketika tipe data dan nilai tidak dituliskan dalam deklarasi konstanta, maka tipe data dan nilai yang dipergunakan adalah sama seperti konstanta yang dideklarasikan diatasnya.
	const (
		a = "konstanta"
		b
	    )
	    fmt.Println(a, b)

	//     deklarasi multiple
	const satu, dua = 1, 2
        const three, four string = "tiga", "empat"
	fmt.Println(satu, dua, three, four)



}

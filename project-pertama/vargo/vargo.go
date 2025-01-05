package vargo

import "fmt"

func Vargo() {
	var firstName string = "Vargo"
	var lastName string = "Palah"
	alias := "mafuyukun"

	sindicate := "wick"
	sindicate = "ethan"
	sindicate = "bourne"

	var nilai1, nilai2, nilai3 string
	nilai1, nilai2, nilai3  = "nilai1", "nilai2", "nilai3"

	// cara lain
	nilai5, nilai6 := "nilai5", "nilai6"
	// multiple variable interface
	one, two, kebenaran, nilaiujian := 1, 2, true, 8.5

	// variable keranjang sampah
	_ = "belajar golang"
	_ = "golang itu mudah"
	name, _ := "kusogakiga", "point"

	// variable pointer
	fullname := new(string)
	fmt.Println(fullname)
	fmt.Println(*fullname)


	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName, "as " + alias + " as " + sindicate)
	fmt.Println("multiple variable 1", nilai1, nilai2, nilai3)
	fmt.Println("multiple variable 2", nilai5, nilai6)
	fmt.Println("multiple variable 3", one, two, kebenaran, nilaiujian, name) //ini ada menggunakan variable keranjang sampah

}

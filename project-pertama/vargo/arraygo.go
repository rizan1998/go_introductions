package vargo

import "fmt"

func Arraygo() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// array tanpa jumlah element
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// jika tambah array tidak bisa langsung
	// Konversi array ke slice
	numbersSlice := numbers[:]

	// Tambah elemen baru
	numbersSlice = append(numbersSlice, 5, 6)

	fmt.Println("data array \t:", numbersSlice)
	fmt.Println("jumlah elemen \t:", len(numbersSlice))

	var h = 0
	var avgFruits = len(fruits)
	for {
		fmt.Println("iterasi ke-", h, "name = ", fruits[h])
		h++
		if h >= avgFruits {
			break
		}
	}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// hanya menampilkan nama buah iterasi tidak di pakai
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// for i, _ := range fruits { }
	// // atau
	// for i := range fruits { }

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var newFruits = fruits[0:2]

	fmt.Println(newFruits) // ["apple", "grape"]

	// Kode	Output	Penjelasan
	// fruits[0:2]	[apple, grape]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
	// fruits[0:4]	[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
	// fruits[0:0]	[]	menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
	// fruits[4:4]	[]	menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
	// fruits[4:0]	[]	error, pada penulisan fruits[a:b] nilai a harus lebih kecil atau sama dengan b
	// fruits[:]	[apple, grape, banana, melon]	semua elemen
	// fruits[2:]	[banana, melon]	semua elemen mulai indeks ke-2
	// fruits[:2]	[apple, grape]	semua elemen hingga sebelum indeks ke-2

	// slice jika di ubah akan impact sampai parent slice
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// perhitungan menggunakan cap tidak akan di update sesuai jumlah elemen

	fmt.Println(fruits)
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aaaFruits = fruits[0:3]
	fmt.Println(aaaFruits)
	fmt.Println(len(aaaFruits)) // len: 3
	fmt.Println(cap(aaaFruits)) // cap: 4

	var bbbFruits = fruits[1:4]
	fmt.Println(bbbFruits)
	fmt.Println(len(bbbFruits)) // len: 3
	fmt.Println(cap(bbbFruits)) // cap: 3

	// var cFruits = append(newFruits, "papaya")
	// fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// var data map[string]int
	// data["one"] = 1
	// // akan muncul error!

	var data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	fmt.Println(chicken1) // map[januari:50 februari:40]
	fmt.Println(chicken2) // map[januari:50 februari:40]

	var _ = map[string]int{}
	var _ = make(map[string]int)
	var _ = *new(map[string]int)

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	var value, isExist = chicken["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

}

// dst := []string{"potato", "potato", "potato"}
// src := []string{"watermelon", "pinnaple"}
// n := copy(dst, src)

// fmt.Println(dst) // watermelon pinnaple potato
// fmt.Println(src) // watermelon pinnaple
// fmt.Println(n)   // 2

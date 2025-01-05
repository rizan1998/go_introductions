package vargo

import "fmt"

func Loopgo(){
	for t := 0; t <10; t++{
		fmt.Printf("t = %d\n", t);
	}


	var l = 0;
	for l < 10{
		println("l = ", l)
		l++;
	}

	var value1 = 0;
	for  {
		fmt.Printf("value1 = %d\n", value1);
		value1++;

		if(value1 == 5){
			break;
		}
	}


	// loop range
	var xs = "123" // string
	for i, v := range xs {
	fmt.Println("Index=", i, "Value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
	fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
	fmt.Println("Value=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
	fmt.Println("Key=", k, "Value=", v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
	fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
	fmt.Print(i) // 01234
	}

	// break and continue
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
		    continue
		}

		if i > 8 {
		    break
		}

		fmt.Println("Angka", i)
	}

	// looping nested
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
		    fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// label for
	outerLoop:
	for i := 0; i < 5; i++ {
	for j := 0; j < 5; j++ {
		if i == 3 {
		break outerLoop
		}
		fmt.Print("matriks [", i, "][", j, "]", "\n")
	}
	}
}

package vargo

import "fmt"

func Vargoif() {
	var point = 0

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}


	var point2 = 8840.0

	if percent := point2 / 100; percent >= 100 {
	fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
	fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
	fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var point3 = 6
	switch point3{
		case 8:
			fmt.Println("perfect")
		case 7, 6, 5, 4:
			fmt.Println("awesome")
		default:
			fmt.Println("not bad")
	}


	var point4 = 6
	switch {
		case point4 == 8:
		fmt.Println("perfect")
		case (point4 < 8) && (point > 3):
		fmt.Println("awesome")
		default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

}

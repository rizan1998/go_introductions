package vargo

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func Funcgo() {
	var randomValue int

	var title string = "belajar golang func"
	fmt.Println(title)

	var names = []string{"john", "wick", "ethan"}
	printMessage("hallo", names)

	randomValue = randomWithRange(1, 100)
	fmt.Println("random value", randomValue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	var dataNilai = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculate2(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var avg2 = calculate2(dataNilai...)

	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	var msg2 = fmt.Sprintf("Rata-rata : %.2f", avg2)
	fmt.Println(msg)
	fmt.Println(msg2)

	yourHobbies("wick", "sleeping", "eating")
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	var numbers2 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var newNumbers2 = func(min int) []int {
		var arr []int
		for _, e := range numbers2 {
			if e < min {
				continue
			}
			arr = append(arr, e)
		}
		return arr
	}(3)

	fmt.Println("original number :", numbers2)
	fmt.Println("filtered number :", newNumbers2)


	var max3 = 3
    var numbers3 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
    var howMany, getNumbers = findMax(numbers3, max3)
    var theNumbers = getNumbers()

    fmt.Println("numbers\t:", numbers)
    fmt.Printf("find \t: %d\n\n", max)

    fmt.Println("found \t:", howMany)    // 9
    fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, ",")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

func calculate2(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

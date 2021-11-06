package main

import (
	"fmt"
	"math"
)

func main() {
	var names [3]string
	names[0] = "Jordan"
	names[1] = "Kate"
	names[2] = "Bob"

	fmt.Println(names)

	array := [3]string{"John", "Kate", "Robert"}
	fmt.Println(array)
	fmt.Println(array, len(array))

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	//***************************************

	marks := [5]float64{11, 9, 12, 8, 10}
	fmt.Println(marks)
	var sum float64 = 0

	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}

	fmt.Println("Cумма чисел: " + fmt.Sprint(sum))

	var result float64 = sum / float64(len(marks))

	fmt.Println(math.Round(result))
}
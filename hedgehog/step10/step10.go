package main

import "fmt"

func main() {
	array := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(array)
	fmt.Println(array[1][2])
	fmt.Println(array[2][3])
}
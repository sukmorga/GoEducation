package main

import "fmt"

func main() {
	defer fmt.Println("Bye")
	fmt.Println("Hello")

	defer fmt.Println("Пример решен")

	var a float64
	var b float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(a + b)
}
package main

import "fmt"

func first() {
	fmt.Println("Hello from first function")
}

func sum(a int, b int) {
	fmt.Println(a + b)
}

func sum2(a int, b int) int {
	result := a + b
	return result
}

func mathFunc(a int, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	return sum, sub, mul, div
}

func mathFunc2(a int, b int) (sum2 int, sub2 int, mul2 int, div2 int) {
	sum2 = a + b
	sub2 = a - b
	mul2 = a * b
	div2 = a / b
	return 
}

func main() {
	first()

	sum(5, 9)

	value := sum2(10, 3)
	fmt.Println(value)

	s, su, m, d  := mathFunc(15, 5)
	fmt.Println(s, su, m, d)

	s2, su2, m2, d2  := mathFunc(15, 5)
	fmt.Println(s2, su2, m2, d2)
}
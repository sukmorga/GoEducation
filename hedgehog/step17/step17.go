package main

import "fmt"

func test(someFunction func(int) int) {
	fmt.Println(someFunction(25))
}

func test2(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	a := func() {
		fmt.Println("Hello")
	}
	a()

	b := func(x int, y int) {
		fmt.Println(x + y)
	}
	b(5, 18)

	c := func(x int, y int) int {
		return x + y
	}
	sum := c(44, 18)
	fmt.Println(sum)

	square := func(x int) int {
		return x * x
	}

	test(square)

	q := test2("Hello")
	q()

	test2("Bye")()
}
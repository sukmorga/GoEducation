package main

import "fmt"

func main() {
	a := 1

	if a >= 1 {
		fmt.Println("a >= 0")
	}

	c := 3
	b := 10
	
	if c > 5 || b > 5 {
		fmt.Println("TRUE")
	}

	v := 5

	if v != 2 {
		fmt.Println("Done!")
	}
}
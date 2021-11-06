package main

import "fmt"

func change(str *string) {
	*str = "LOL"
}

func main() {
	a := 5
	pointer := &a
	fmt.Println(pointer)
	fmt.Println(*pointer)

	s := "LLL"
	fmt.Println(s)
	change(&s)
	fmt.Println(s)

	q := "aaaa"
	fmt.Println(q)
	var pointer2 *string = &q
	change(pointer2)
	fmt.Println(s)

	num := 9
	b := &num
	fmt.Println(*b)
	*b = 35
	fmt.Println(*b)
}
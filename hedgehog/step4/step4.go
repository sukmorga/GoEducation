package main

import "fmt"

func main() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}

	var i int = 0
	for ; i <= 50; i += 10 {
		fmt.Println(i)
	}

	var j int = 0
	for j <= 50 {
		fmt.Println(j)
		j += 10
	}
}

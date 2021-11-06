package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 1, 2, 5, 7, 4}
	fmt.Println(slice)

	slice = append(slice, 0)
	fmt.Println(slice)

	slice[0] = 11
	fmt.Println(slice)

	sort.Ints(slice)
	fmt.Println(slice)

	slice2 := []string{"v", "b", "a", "c", "d", "z"}
	slice2 = append(slice2, "hhh")
	fmt.Println(slice2)
	sort.Strings(slice2)
	fmt.Println(slice2)
}
package main

import "fmt"

func main() {
	var money map[string]int = map[string]int{
		"dollars": 1000,
		"euros":   2000,
	}
	fmt.Println(money)
	fmt.Println(money["dollars"])

	money2 := map[string]int{
		"dollars": 30000,
		"euros": 4367,
	}
	fmt.Println(money2)

	money["dollars"] = 5000
	fmt.Println(money)

	delete(money, "euros")
	fmt.Println(money)

	fmt.Println(money["gggg"])

	el, status := money["dollars"]
	fmt.Println(el, status)

	el1, status1 := money["fgffg"]
	fmt.Println(el1, status1)
}
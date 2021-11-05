package main

import "fmt"

func main() {
	var a int = 27
	fmt.Println(a)

 	var number float64 = 27.45454
	fmt.Println(number)

	name := "Денис"
	fmt.Println(name)
	fmt.Println(len(name))

	var name2 string
	var age uint
	fmt.Println("Whats is your name?")
	fmt.Scan(&name2)
	fmt.Println("Hello " + name2 + "!")
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	fmt.Println("You are " + fmt.Sprint(age) + " years")

	var h int8 = 2
	var g int64 = int64(h)
	fmt.Println(g)

	//********************************************************LESSON2

}

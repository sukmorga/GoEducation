package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 3
	var b float64 = 5
	sum := a / b
	fmt.Println(sum)


	var c float64 = 144
	result := math.Sqrt(c)
	fmt.Println(result)


	var d float64 = 25.555656565
	result2 := math.Ceil(d)
	fmt.Println(result2)


	var e float64 = 45.53233434
	result3 := math.Floor(e)
	fmt.Println(result3)


	var f float64 = 33.5656565
	result4 := math.Round(f)
	fmt.Println(result4)


	fmt.Println("Дебильный калькулятор")
	fmt.Println("Какое действие вы хотите выполнить? (+, -, *, /)")
	
	var action string
	var number1 float64
	var number2 float64

	fmt.Scan(&action)
	fmt.Println("Введите первое число")
	fmt.Scan(&number1)
	fmt.Println("Введите второе число")
	fmt.Scan(&number2)

	switch {
	case action == "+":
		fmt.Println("Сумма чисел " + fmt.Sprint(number1 + number2))
	case action == "-":
		fmt.Println("Разность чисел " + fmt.Sprint(number1 - number2))
	case action == "*":
		fmt.Println("Произведение чисел " + fmt.Sprint(number1 * number2))
	case action == "/":
		fmt.Println("Результат деления " + fmt.Sprint(number1 / number2))
	default:
		fmt.Println("Калькулятор не выполняет таких операций")
	}
}
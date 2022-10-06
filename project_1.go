package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Print("Введите любое число: ")
	fmt.Scan(&a)
	b = a * 2
	c = b + 100

	fmt.Print(c)
}

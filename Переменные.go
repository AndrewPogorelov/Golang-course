package main

import "fmt"

func main() {

	var Hello string

	Hello = "Hello?"

	var a int = 2019

	fmt.Println(Hello)
	fmt.Println(a)

	// Можно группировать:
	var (
		x int     = 10
		c string  = "Что бы спросить?"
		z float64 = 1.045
	)
	fmt.Println(x, c, z)

	// А еще можно записать так:
	var symbol int32 = '\u7777'
	var symbol1 uint64 = '\u5555' // будет какая нибудь фигня в виде иероглифа
	fmt.Println(string(symbol) + "" + string(symbol1))

	var d int = 5
	fmt.Println(d)
	d--
	fmt.Println(d)
}

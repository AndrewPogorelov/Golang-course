// Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	b = a % 100
	c = b / 10

	fmt.Print(c)
}

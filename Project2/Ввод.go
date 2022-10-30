package main

import "fmt"

func main() {
	var age int
	var name string

	fmt.Print("Введите ваш возраст:")
	fmt.Scan(&age)

	fmt.Print("Введите своё имя:")
	fmt.Scan(&name)

	fmt.Println(age, name)
}

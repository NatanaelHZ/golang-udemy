package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))
	fmt.Println("Teste")

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado - concatena com valor de 123 da ascii
	fmt.Println("Teste " + string(123))

	// int to string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string to int
	num, err := strconv.Atoi("123")
	fmt.Println(num-122, err)

	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	if b {
		fmt.Println("Verdadeiro")
	}
}

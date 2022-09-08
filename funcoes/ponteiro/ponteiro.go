package main

import "fmt"

func inc1(n int) {
	n++
}

// Revisão ponteiro é representado por *
func inc2(n *int) {
	fmt.Println(n)
	fmt.Println(&n)

	// Revisão: é usado para desrefenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)
	fmt.Println(n)

	// Revisão: & usado para obter o endereço da váriavel
	inc2(&n)
	fmt.Println(n)
}

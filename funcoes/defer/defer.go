package main

import "fmt"

func obterAprovado(numero int) int {
	defer fmt.Println("Fim!") // Executa default antes do return
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterAprovado(6000))
	fmt.Println()
	fmt.Println(obterAprovado(3000))
}

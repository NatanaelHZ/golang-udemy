package main

import "fmt"

func main() {
	numeros := [...]int{7, 2, 3, 4, 5}

	fmt.Println("")
	for i, numero := range numeros {
		fmt.Printf("[%d] => %d \n", i, numero)
	}

	fmt.Println("")
	for _, numero := range numeros { //Ignorar indice
		fmt.Printf("=> %d \n", numero)
	}
}

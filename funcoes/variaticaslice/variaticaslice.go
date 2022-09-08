package main

import "fmt"

func imprimirAprovador(aprovados ...string) {
	fmt.Printf("Lista de Aprovados:\n\n")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s \n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Gandalf", "Elrond", "Galadriel", "Sam"}

	imprimirAprovador(aprovados...)
}

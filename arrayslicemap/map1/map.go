package main

import "fmt"

func main() {
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Gandalf"
	aprovados[987654321] = "Elrond"
	aprovados[789456123] = "Galadriel"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("Nome: %s CPF: %d\n", nome, cpf)
	}

	fmt.Println("Aprovado: ", aprovados[123456789])
	delete(aprovados, 123456789)

	fmt.Println(aprovados)
}

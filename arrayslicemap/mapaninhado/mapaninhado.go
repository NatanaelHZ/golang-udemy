package main

import "fmt"

func main() {
	funcsPorletra := map[string]map[string]float64{
		"G": {
			"Galadriel": 10000.00,
			"Gandalf":   10000.00,
		},
		"E": {
			"Elrond": 9900.00,
		},
		"T": {
			"Telden": 8000.00,
		},
	}

	fmt.Println(funcsPorletra)

	for letra, funcionarios := range funcsPorletra {
		fmt.Printf("\n\nLetra: %s\n", letra)

		for nome, salario := range funcionarios {
			fmt.Printf("Nome: %s Salário: %.2f", nome, salario)
		}
	}

	fmt.Printf("\n Deletando \n")
	delete(funcsPorletra, "G")

	for letra, funcionarios := range funcsPorletra {
		fmt.Printf("\n\nLetra: %s\n", letra)

		for nome, salario := range funcionarios {
			fmt.Printf("Nome: %s Salário: %.2f", nome, salario)
		}
	}
}

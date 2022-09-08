package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Elrond":    11000.00,
		"Galadriel": 12000.00,
		"Gandalf":   13000.00,
	}

	fmt.Println(funcsESalarios)

	funcsESalarios["Radagast"] = 9000.00

	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Printf("Nome: %s | Salário: %.2f\n", nome, salario)
	}
}

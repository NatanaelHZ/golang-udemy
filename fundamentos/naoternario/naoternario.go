package main

import "fmt"

//Não tem operador ternario

func obterResultadoNota(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(obterResultadoNota(5))
}

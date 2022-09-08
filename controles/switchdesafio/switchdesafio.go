package main

import "fmt"

func notaParaConceito(nota float64) string {

	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	var nota float64

	nota = 2
	fmt.Println(notaParaConceito(nota))
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(8.5))
	fmt.Println(notaParaConceito(6.5))
	fmt.Println(notaParaConceito(4.5))
	fmt.Println(notaParaConceito(2.5))
}

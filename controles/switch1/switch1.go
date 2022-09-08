package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota int64 = int64(n)

	switch nota {
	case 10:
		fallthrough //Continua execução
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(11))
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(8.5))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(1))
	fmt.Println(notaParaConceito(-2))
}

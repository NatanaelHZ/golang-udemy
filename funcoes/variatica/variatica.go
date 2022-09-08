package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f \n", media(2, 10, 10))
	fmt.Printf("Média: %.2f \n", media(9, 10, 10, 10))
}

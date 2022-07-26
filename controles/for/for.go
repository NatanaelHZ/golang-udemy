package main

import (
	"fmt"
	"time"
)

// Não existe while
func main() {
	i := 1

	//Semelhante ao while
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	for {
		//Laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 5)
	}
}

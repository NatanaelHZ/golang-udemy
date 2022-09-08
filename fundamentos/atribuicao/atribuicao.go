package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //inferência de tipo
	fmt.Println(i)
	i += 3 // i = i + 3
	fmt.Println(i)
	i -= 3 // i = i - 3
	fmt.Println(i)
	i /= 3 // i = i / 2
	fmt.Println(i)
	i *= 3 // i = i * 2
	fmt.Println(i)
	i %= 3 // i = i % 2
	fmt.Println(i)

	// Atribuição multipla
	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

}

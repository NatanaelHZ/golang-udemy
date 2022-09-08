package main

import "fmt"

func main() {
	numeros := [...]int{7, 2, 3, 4, 5}
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7

	fmt.Println(s1, s2)

	s3 := numeros[2:4]
	fmt.Println(s3)
	s3[0] = 8
	fmt.Println(s3)
	fmt.Println(numeros)

}

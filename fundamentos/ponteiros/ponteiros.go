package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // Atruindo o endereço da variavel
	*p++   // Desreferencinado
	i++
	fmt.Println(p, *p, i, &i)

	//Go não tem aritimética de ponteiro
	//p++
}

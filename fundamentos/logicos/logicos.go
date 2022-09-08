package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2
	tomarSorvete := trab1 || trab2

	return comprarTV50, comprarTV32, tomarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)

	fmt.Printf("tv 50: %t, tv 32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}

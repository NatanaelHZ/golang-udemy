package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// Pode adicionar mais métodos
}

type bmwS7 struct{}

func (b bmwS7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func (b bmwS7) ligarTurbo() {
	fmt.Println("Ligar turbo")
}

func main() {
	var b esportivoLuxuoso = bmwS7{}
	b.ligarTurbo()
	b.fazerBaliza()
}

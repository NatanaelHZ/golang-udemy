package main

import "math"

// Inicializando com letra maiúscula é PÚBLICO (visibilidade dentro e fora do pacote)!
// Inicializando com letra minúscula é PRIVADO (visivel apenas dentro do pacote)

// Por exemplo
// Ponto - gerará algo público
// ponto ou _ponto - gerará algo privado

// Ponto representa uma cordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return
}

// Distância é responsavel por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

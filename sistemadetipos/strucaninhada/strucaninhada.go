package main

import "fmt"

type item struct {
	produtoId int
	qtd       int
	preco     float64
}

type pedido struct {
	userId int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}

	return total
}

func main() {
	pedido := pedido{
		userId: 1,
		itens: []item{
			item{1, 2, 12}, //Atribuindo pelo indice
			item{
				produtoId: 21,
				qtd:       10,
				preco:     50,
			},
		},
	}

	fmt.Printf("Valor total pedido: %.2f", pedido.valorTotal())
}

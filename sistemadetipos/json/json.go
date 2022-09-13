package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct para json

	p1 := produto{1, "Livro", 50.0, []string{"Promoção", "Cultura"}}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	// Json para string
	var p2 produto
	jsonString := `{"id":2, "nome":"Bola", "preco":50.0, "tags":["Promoção", "Lazer"]}`

	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)
}

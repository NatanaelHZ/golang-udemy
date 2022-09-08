package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAlatorio() int {
	s := rand.NewSource(time.Now().UnixNano())

	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	if i := numeroAlatorio(); i > 5 { // tbm supórtado no switch
		fmt.Println("Ganhou", i)
	} else {
		fmt.Println("Perdeu", i)
	}
}

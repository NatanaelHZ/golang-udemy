package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	h := t.Hour()

	switch { //switch true
	case h < 12:
		fmt.Println("Bom dia!", h)
	case h < 18:
		fmt.Println("Boa tarde!", h)
	default:
		fmt.Println("Boa noite!", h)
	}
}

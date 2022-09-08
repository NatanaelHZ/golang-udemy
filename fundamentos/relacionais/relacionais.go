package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("3 != 2:", 3 != 2)
	fmt.Println("3 < 2:", 3 < 2)
	fmt.Println("3 > 2:", 3 > 2)
	fmt.Println("3 <= 2:", 3 <= 2)
	fmt.Println("3 >= 2:", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println("Data: ", d1 == d2)
	fmt.Println("Datas: ", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

}

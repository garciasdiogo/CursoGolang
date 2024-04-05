package main

import "fmt"

func main() {
	idade := 20

	if idade >= 18 {
		fmt.Println("Maior de idade")

	} else {
		fmt.Println("Menor de idade")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

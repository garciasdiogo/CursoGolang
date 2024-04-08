package main

import "fmt"

func main() {

	// Estrutura de controle: if
	idade := 18
	if idade >= 18 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}

	//Estrutura de controle: for
	for i := 0; i < 5; i++ {
		fmt.Println("Número:", i)
	}

	//Estrutura de controle: switch
	dia := "sábado"
	switch dia {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("Dia útil")
	case "sábado", "domingo":
		fmt.Println("Fim de semana")
	default:
		fmt.Println("Dia inválido")
	}

	/*
		if é usado para fazer verificações condicionais.
		for é usado para loops.
		switch é uma forma concisa de escrever uma série de instruções if-else.
	*/
}

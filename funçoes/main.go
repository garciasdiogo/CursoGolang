package main

import "fmt"

//Função que retorna a soma de dois números inteiros
func soma(a, b int) int {
	return a + b
}

//Função com múltiplos valores de retorno
func divisao(a, b int) (int, int) {
	quociente := a / b
	resto := a % b
	return quociente, resto
}

func main() {
	// Chamando funções
	resultado := soma(3, 5)
	fmt.Println("Soma:", resultado)

	q, r := divisao(10, 3)
	fmt.Println("Quociente:", q, "Resto", r)
}

//Funções são blocos de construção fundamentais em Go.
//Os parâmetros são seguidos pelo tipo.
//Você pode retornar múltiplos valores de uma função.

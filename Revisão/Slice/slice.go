package main

import (
	"fmt"
)

func main() {
	exibeNomes()
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "João"}
	// Len = Quantidade de itens no vetor
	fmt.Println("O meu slice tem", len(nomes), "itens")
	// Cap = Capacidade no vetor
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	//Append = adicionar item
	nomes = append(nomes, "Aparecida")
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}

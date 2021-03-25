package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	fmt.Println("Menu: ")
	fmt.Println("1- Teste 1")
	fmt.Println("2- Teste 2")
	fmt.Println("0- Sair do Programa")
	fmt.Print("Qual comando deseja executar: ")

	comando := leComando()
	switch comando {
	case 1:
		fmt.Println("Teste 1")
	case 2:
		fmt.Println("Teste 2")
	case 0:
		fmt.Println("Saiu do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando!!")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Kelly"
	fmt.Println("Olá, sr(a).", nome)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

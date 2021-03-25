package main

import "fmt"

func main() {

	var comando int

	fmt.Println("1- Teste 1")
	fmt.Println("2- Teste 2")
	fmt.Println("0- Sair do Programa")
	fmt.Print("Qual comando deseja executar: ")
	fmt.Scan(&comando)

	/*if comando == 1 {
		fmt.Println("Teste 1")
	} else if comando == 2 {
		fmt.Println("Teste 2")
	} else if comando == 3 {
		fmt.Println("Saida")
	} else {
		fmt.Println("Não reconheço este comando!")
	}*/

	switch comando {
	case 1:
		fmt.Println("Teste 1")
	case 2:
		fmt.Println("Teste 2")
	case 3:
		fmt.Println("Saida")
	default:
		fmt.Println("Não reconheço este comando!")
	}

}

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	//	_, idade := devolveNomeIdade()
	//	fmt.Println(idade)
	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		//	os.Exit(-1)
	}
}

/* func devolveNomeIdade() (string, int) {
	nome := "Kelly"
	idade := 23
	return nome, idade
}*/
func exibeIntroducao() {
	nome := "Kelly"
	fmt.Println("Olá, sr(a).", nome)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://posgraduacao.mackenzie.br/login/index.php"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}

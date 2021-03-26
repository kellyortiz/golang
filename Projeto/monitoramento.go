package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	// "while"
	for {
		exibeMenu()

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
			os.Exit(-1)
		}
	}
}
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
	var sites [4]string
	sites[0] = "https://posgraduacao.mackenzie.br/login/index.php"
	sites[1] = "https://github.com/"
	sites[2] = "https://www.linkedin.com/feed/"
	fmt.Println(sites)

	site := "https://posgraduacao.mackenzie.br/login/index.php"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, " foi carregado com sucesso!!")
	} else {
		fmt.Println("Site:", site, " está com problemas. Status Code:", resp.StatusCode)
	}
}

/* Anotação:
		_, idade := devolveNomeIdade()
        fmt.Println(idade)
		func devolveNomeIdade() (string, int) {
		nome := "Kelly"
		idade := 23
		return nome, idade
		fmt.Println(resp)
*/

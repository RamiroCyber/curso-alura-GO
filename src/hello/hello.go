package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const tempoDeEspera = 5

func main() {
	welcome()
	for {
		exibirMenu()
		comando := opcaoComando()

		switch comando {
		case 1:
			inciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}
func welcome() {
	nome := "Ramiro"
	versao := 1.2
	fmt.Println("Olá", nome)
	fmt.Println("Sistema com a versäo", versao)
}

func exibirMenu() {
	fmt.Println("1- Inciar monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func opcaoComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi ", comandoLido)
	return comandoLido
}

func inciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lerSitesArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			fmt.Println("Testando site", site)
			testeSite(site)
		}
		time.Sleep(tempoDeEspera * time.Second)
	}

}

func testeSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro")
	}
	fmt.Println(string(arquivo))
}

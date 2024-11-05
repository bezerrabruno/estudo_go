package monitoramento

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"scripts/logs"
	"strings"
	"time"
)

func IniciarMonitoramento() {
	for {
		fmt.Println("Monitorando...")
		sites := leSitesDoArquivo()

		for _, site := range sites {
			fmt.Println("")
			testaSite(site)
		}
		time.Sleep(5 * time.Second)
		fmt.Println("")
		fmt.Println("")
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		logs.RegistraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		logs.RegistraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("db/sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
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

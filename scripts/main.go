package main

import (
	"fmt"
	"os"
	"scripts/logs"
	"scripts/mensagens"
	"scripts/monitoramento"
	"scripts/scans"
)

func main() {
	mensagens.Introducao()

	mensagens.Menu()

	comando := scans.ComandoEscolhido()

	switch comando {
	case 1:
		monitoramento.IniciarMonitoramento()

	case 2:
		logs.ImprimeLogs()

	case 0:
		finalizarPrograma(0)

	default:

		finalizarPrograma(-1)
	}

}

func finalizarPrograma(id int) {
	switch id {
	case -1:
		fmt.Println("Comando não reconhecido")
		os.Exit(-1)

	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)

	default:
		fmt.Println("Algo inesperado aconteceu, tente novamente!")
		os.Exit(-1)
	}

}

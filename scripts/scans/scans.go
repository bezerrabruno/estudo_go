package scans

import "fmt"

func ComandoEscolhido() int {
	var comandoLido int
	_, err := fmt.Scan(&comandoLido)

	if err != nil {
		return 0
	}

	fmt.Println("")

	return comandoLido
}

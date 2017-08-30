package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	exibirIntroducao()
	for {
		exibirMenu()
		comando := leComando()
		switch comando {
		case 1:
			if err := quantoFaltaPS(); err != nil {
				fmt.Println(err)
			}
		case 0:
			os.Exit(0)
		}

	}
}
func exibirIntroducao() {
	version := 1.0
	fmt.Println("---------------------------------------------------------")
	fmt.Println("	Bem vindo ao Quanto falta para a PS - version", version)
	fmt.Println("---------------------------------------------------------")
	fmt.Println("")
}

func exibirMenu() {
	fmt.Println("O que deseja fazer?")
	fmt.Println("1 - Calcular o quanto falta para a PS")
	fmt.Println("0 - Sair do programa")
}
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}
func quantoFaltaPS() (err error) {
	var nac float32
	fmt.Println("Digite a média final de NAC: ")
	fmt.Scan(&nac)
	var am float32
	fmt.Println("Digite a média final de AM: ")
	fmt.Scan(&am)

	if nac > 10 || nac < 0 || am > 10 || am < 0 {
		err = errors.New("notas Inválidas")
		return
	}
	ps := calcular(nac, am)
	if ps > 10 {
		fmt.Println("Você precisa de um milagre pra passar! Desculpe :(")
		fmt.Println("")
		return
	}
	fmt.Printf("\nVocê precisa tirar %0.2f na PS para passar de ano!\n", ps)
	return
}

func calcular(nac float32, am float32) float32 {
	media := nac*0.2 + am*0.3
	ps := (6 - media) * 2
	return ps
}

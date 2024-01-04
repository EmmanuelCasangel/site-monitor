package main

import (
	"fmt"
	"os"
)

func main() {

	displaysIntroduction()
	displaysMenu()
	comand := readComand()

	switch comand {
	case 1:
		fmt.Println("Iniciando monitoramento")
	case 2:
		fmt.Println("Exibindo logs:")
	case 0:
		fmt.Println("Até uma proxima :)")
		os.Exit(0)
	default:
		fmt.Println("Essa opcao nao esta listada")
		os.Exit(-1)
	}
}

func displaysIntroduction() {
	name := "Emmanuel"
	var version = 1.2

	fmt.Println("Oie, fala", name)
	fmt.Println("Esse programa esta na versão ", version)
}

func displaysMenu() {
	fmt.Println("1- Iniciar monitoramento de sites")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func readComand() int {
	var comand int
	fmt.Scan(&comand)
	fmt.Println("O comando escolhido foi:", comand)

	return comand
}

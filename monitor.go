package main

import "fmt"

func main() {
	nome := "Emmanuel"
	var version = 1.2

	fmt.Println("Oie, fala", nome)
	fmt.Println("Esse programa esta na versão ", version)

	fmt.Println("1- Iniciar monitoramento de sites")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)

	if comando == 1 {
		fmt.Println("Iniciando monitoramento")

	} else if comando == 2 {
		fmt.Println("Exibindo logs:")

	} else if comando == 3 {
		fmt.Println("Até uma proxima :)")

	} else {
		fmt.Println("Essa opcao nao esta listada")
	}

}

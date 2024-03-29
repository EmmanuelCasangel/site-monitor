package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 5 * time.Second
const urlsFileName = "urls.txt"
const logsFileName = "logs.txt"
const monitoringQuantity = 2

func main() {
	displaysIntroduction()

	for {
		displaysMenu()
		comand := readComand()

		switch comand {
		case 1:
			fmt.Println("Iniciando monitoramento")
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs:")
			displayLogs()
		case 0:
			fmt.Println("Até uma proxima :)")
			os.Exit(0)
		default:
			fmt.Println("Essa opcao nao esta listada")
			os.Exit(-1)
		}
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

	return comand
}

func startMonitoring() {
	// urls := []string{
	// 	"http://www.alura.com.br",
	// 	"https://httpbin.org/status/500",
	// 	"https://www.caelum.com.br",
	// 	"https://github.com/EmmanuelCasangel/site-monitor",
	// }

	urls := readUrlsFromFile(urlsFileName)

	for i := 0; i < monitoringQuantity; i++ {
		fmt.Println()
		fmt.Println("Realizando monitoramentos:")

		for _, url := range urls {

			verifyUrl(url)
		}
		time.Sleep(delay)
	}

}

func verifyUrl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", url, "foi carregado com sucesso")
		recordLogs(url, true)
	} else {
		fmt.Println("Site", url, "esta com problemas Status Code:", resp.StatusCode)
		recordLogs(url, false)
	}
}

func readUrlsFromFile(fileName string) []string {
	var urls []string

	file, err := os.Open(fileName)
	// file, err := os.ReadFile("urls.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		urls = append(urls, line)

		if err == io.EOF {
			break
		}

	}

	file.Close()

	return urls
}

func recordLogs(url string, status bool) {
	file, err := os.OpenFile(
		logsFileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0666,
	)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)

	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func displayLogs() {
	file, err := os.ReadFile(logsFileName)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(string(file))

}

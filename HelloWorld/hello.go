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

const monitoring_times = 3
const delay = 5

func main(){
	readWebsitesFromFile()
	intro()
	
	command := getCommand()

	switch command {
	case 1:
		fmt.Println("monitoring")
		startMonitoring()
	case 2:
		fmt.Println("logging")
		printLogs()
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Comando não existe")
		os.Exit(-1)
	}
}

func intro(){
	name := "Xablau"
	version := 1.01

	fmt.Println("Olá, senhor", name)
	fmt.Println("ESte programa está na versão", version)
}

func getCommand() int {
	var command int

	fmt.Println("1 - Iniciar o monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	fmt.Scan(&command)

	fmt.Println("O valor da variável comando é:", command)

	return command
}

func startMonitoring(){
	websites := readWebsitesFromFile()
	
	for i:=0; i<monitoring_times; i++{
		for _, website := range(websites){
			testWebsiteHealth(website)
		}
		time.Sleep(delay*time.Second)
	}
	
}

func testWebsiteHealth(website string){
	response, err := http.Get(website)

	if response.StatusCode == 200{
		fmt.Println("Site", website, "foi carregado com sucesso")
		recordLog(website, true)
	} else {
		fmt.Println("Site", website, "obteve erros:", err, "status: ", response.StatusCode)
		recordLog(website, false)
	}
}

func readWebsitesFromFile() []string{
	var websites []string

    file, err := os.Open("HelloWorld/websites.txt")
    
	if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    reader := bufio.NewReader(file)

    for {
        line, err := reader.ReadString('\n')
        line = strings.TrimSpace(line)
        websites = append(websites, line)
        if err == io.EOF {
            break
        }
    }

    file.Close()

    return websites
}

func recordLog(site string, status bool) {

    file, err := os.OpenFile("HelloWorld/log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    file.WriteString(time.Now().Format("01/02/2006 15:04:05") + " - "+ site + "- online: " + strconv.FormatBool(status) + "\n")

    file.Close()
}

func printLogs(){
	file, err := ioutil.ReadFile("HelloWorld/log.txt")

	if err == nil{ 
		fmt.Println(err)
	}
	fmt.Println(string(file))

}
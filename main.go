package main

import(
"os"
//"os/exec"
"io"
"io/ioutil"
"encoding/json"
"fmt"
"bytes"
)
var out bytes.Buffer
var exit string
var option int
var answerQuit = true
var addServer string
var nameServer string
var storeFileName = "servers.json"
var optionServer int
type ServerData struct{
	Name string
	Ip string
	Login string
	Password string
}


func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func checkStatus(status bool) bool{
	if status {
		return true
	}
	return false
}

func main() {
	arrayServers := make([]ServerData,0)
	TotalServersAdded := 1
	for{
		fmt.Println("----------- Digite umas das opções abaixo -----------")
		fmt.Println("1 -> Cadastrar novo servidor na lista de servidores")
		fmt.Println("2 -> Ver dados de um determinado Servidor")
		fmt.Println("3 -> Conectar a um servidor via ssh")
		fmt.Println("4 -> Apagar todos servidores cadastrados")
		fmt.Scan(&option)

		switch(option){
		case 1:
			parseServers(&arrayServers)
			for{
				data := ServerData{}
				addServerToArray(&data)
				arrayServers = append(arrayServers,data)
				fmt.Println("Deseja adicionar mais um servidor? (Y/N)")
				fmt.Scan(&addServer)
				if addServer == "N" {
					break
				}
				TotalServersAdded++
			}
			store(&arrayServers,&TotalServersAdded,false)
			break;
		case 2:
			parseServers(&arrayServers)
			listServers(arrayServers)
			fmt.Println("Selecione pela opção: ( -1 para sair das opções e voltar pra tela inicial)")
			serverSelect, ok := selectServer(arrayServers)
			if checkStatus(ok) {
				showDataServer(serverSelect)
			}else{
				answerQuit = false
			}
			break;
		case 3:
			parseServers(&arrayServers)
			listServers(arrayServers)
			fmt.Println("Selecione pela opção: ( -1 para sair das opções e voltar pra tela inicial)")
			serverSelect, ok := selectServer(arrayServers)
			if checkStatus(ok) {
				fmt.Printf("ssh %s@%s -P %s \n",serverSelect.Login,serverSelect.Ip,serverSelect.Password)
			}else{
				answerQuit = false
			}
			break;
		case 4:
			os.Truncate(storeFileName,1)
			store(&arrayServers,&TotalServersAdded,true)
			break;
		}

		if answerQuit {
			fmt.Println("Deseja sair do programa? (Y/N)")
			fmt.Scan(&exit)
			if exit == "Y" {
				break
			}
		}
	}
}

func parseServers(arrayServers *[]ServerData) {
	servers, err := ioutil.ReadFile(storeFileName)
	checkError(err)
	json.Unmarshal(servers, arrayServers)
}

func addServerToArray(data *ServerData) bool{
	fmt.Println("Nome Servidor: ")
	fmt.Scan(&data.Name)
	fmt.Println("Ip Servidor")
	fmt.Scan(&data.Ip)
	fmt.Println("Usuário do servidor")
	fmt.Scan(&data.Login)
	fmt.Println("Senha do servidor")
	fmt.Scan(&data.Password)
	return true
}

func listServers(arrayServers []ServerData){
	count := 1
	for _, v := range arrayServers{
		fmt.Println(count,")", v.Name)
		count++
	}
}

func selectServer(arrayServers []ServerData) (ServerData, bool){
	fmt.Scan(&optionServer)
	if optionServer == -1 {
		return ServerData{},false
	}
	return arrayServers[optionServer-1],true
}

func showDataServer(serverSelect ServerData){
	fmt.Println("")
	fmt.Println("-------------------------------")
	fmt.Println("Server:",serverSelect.Name)
	fmt.Println("Ip:",serverSelect.Ip)
	fmt.Println("Usuário:",serverSelect.Login)
	fmt.Println("Senha:",serverSelect.Password)
	fmt.Println("-------------------------------")
	fmt.Println("")
}

func store(arrayServers *[]ServerData,TotalServersAdded *int,notServersAdded bool) {
	jencoded, err := json.Marshal(*arrayServers)
	checkError(err)
	file, err := os.OpenFile(storeFileName,os.O_CREATE|os.O_RDWR, 0666)
	checkError(err)
	buffer , err := io.WriteString(file,string(jencoded))
	checkError(err)
	file.Close()

	if notServersAdded {
		fmt.Println("Todos os servidores foram apagados com sucesso!")
	}else{
		if buffer > 0 {
			if *TotalServersAdded == 1 {
				fmt.Println("Servidor adicionado com sucesso!")
			}else{
				fmt.Println("Servidores adicionados com sucesso!")
			}
		}
	}
}
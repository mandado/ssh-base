package main

import(
"os"
"io"
"io/ioutil"
"encoding/json"
"fmt"
"bytes"
)
var out bytes.Buffer
var exit string
var choice int
var addServer string
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

func main() {
	arrayServers := make([]ServerData,0)
	TotalServersAdded := 1
	for{
		fmt.Println("----------- Digite umas das opções abaixo -----------")
		fmt.Println("1 -> Cadastrar novo servidor na lista de servidores")
		fmt.Scan(&choice)

		if choice == 1 {
			servers, err := ioutil.ReadFile("servers.json")
			checkError(err)
			json.Unmarshal(servers, &arrayServers)
			for{
				data := ServerData{}
				fmt.Println("Nome Servidor: ")
				fmt.Scan(&data.Name)
				fmt.Println("Ip Servidor")
				fmt.Scan(&data.Ip)
				fmt.Println("Usuário do servidor")
				fmt.Scan(&data.Login)
				fmt.Println("Senha do servidor")
				fmt.Scan(&data.Password)
				arrayServers = append(arrayServers,data)
				fmt.Println("Deseja adicionar mais um servidor? (Y/N)")
				fmt.Scan(&addServer)
				if addServer == "N" {
					break
				}
				TotalServersAdded++
			}
		}

		jencoded, err := json.Marshal(arrayServers)
		checkError(err)
		file, err := os.OpenFile("servers.json",os.O_CREATE|os.O_RDWR, 0666)
		checkError(err)
		buffer , err := io.WriteString(file,string(jencoded))
		checkError(err)
		file.Close()
		if buffer > 0 {
			if TotalServersAdded == 1 {
				fmt.Println("Servidor adicionado com sucesso!")
			}else{
				fmt.Println("Servidores adicionados com sucesso!")
			}
		}


		fmt.Println("Deseja Sair? (Y/N)")
		fmt.Scan(&exit)
		if exit == "Y" {
			break
		}
	}
}
package main

import(
"fmt"
"os"
)

func main() {
	exit := true
	var choiceTask int

	if len(os.Args) == 1 {
		fmt.Println("Uso: ex: 1")
		os.Exit(1)
	}

	fmt.Printf("Uso: ex: %d",choiceTask)


	for(exit == true) {
		fmt.Println("----------- Digite umas das opções abaixo -----------")
		fmt.Println("1 -> Cadastrar novo servidor na lista de servidores")
		fmt.Println("2 -> Pesquisar e acessar um servidor via ssh")
		fmt.Scan(&choiceTask)
	}

}
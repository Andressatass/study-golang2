package main

import "fmt"

// geralmente usamos canais em funções separadas
func main() {
	canal := make(chan string, 2) //buffer é definir o tamanho(?)

	canal <- "Ola mundo"
	canal <- "Programando em go"

	mensagem := <-canal
	fmt.Print(mensagem)
}

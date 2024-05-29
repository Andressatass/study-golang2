package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("ola imundo", canal)

	// for {
	// 	mensagem, aberto := <-canal //esperando receber um valor
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // enviando texto para o canal
		time.Sleep(time.Second)
	}

	close(canal)
}

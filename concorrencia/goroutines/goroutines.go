package main

import (
	"fmt"
	"time"
)

//concorrencia != paralelismo

func main() {
	go escrever("ola mundo") // goroutine
	escrever("programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

package main

import "fmt"

func funcao1() {
	fmt.Println("func 1")
}

func funcao2() {
	fmt.Println("func 2")
}

func main() {
	defer funcao1()
	//defer = adiar
	funcao2()
}

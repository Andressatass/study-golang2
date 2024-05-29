package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"
	fmt.Println(texto)

	funcao := func() {
		texto = "dentro da func"
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}

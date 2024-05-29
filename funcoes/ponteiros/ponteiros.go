package main

import "fmt"

func inverteSInal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
	// usa o *x para pegar o valor de um endereço &x
}

func main() {
	numero := 20
	numeroInvertido := inverteSInal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) //& significa enviar o endereço
	fmt.Println(novoNumero)

}

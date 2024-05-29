package main

import "fmt"

func calcula(n1, n2 int) (soma int, subtraao int) {
	soma = n1 + n2
	subtraao = n1 - n2
	return
}

func main() {
	soma, subtracao := calcula(10, 20)
	fmt.Println(soma, subtracao)
}

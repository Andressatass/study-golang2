package main

import "fmt"

func main() {
	fmt.Println("ponteiros")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	fmt.Println("---------------------------")
	// PONTEIRO é uma referencia de memoria
	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3 //guardando o endereço da var3

	fmt.Print(var3, *ponteiro) //desreferenciação
}

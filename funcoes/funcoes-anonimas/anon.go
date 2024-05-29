package main

import "fmt"

func main() {

	retorno := func(test string) {
		fmt.Println(test)
		return fmt.Print("")
	}("passando parametro")

}

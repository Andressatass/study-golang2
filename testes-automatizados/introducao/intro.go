package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoDeEndereco := enderecos.TipoDeEndereco("xdAvenida paulista")
	fmt.Println(TipoDeEndereco)
}

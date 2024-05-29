package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoMinusculo, " ")[0]

	enderecoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return primeiraPalavraEndereco
	}

	return "tipo invalido"
}

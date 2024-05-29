package enderecos

import "testing"

type cenarioDeTeste struct {
	endInserido string
	endEsperado string
}

// teste de unidade
func TestTipoEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua abc", "rua"},
		{"avenida abc", "avenida"},
		{"rodovia abc", "rodovia"},
		{"Praça abc", "tipo invalido"},
		{"estrada abc", "estrada"},
	}

	for _, cenario := range cenariosDeTeste {
		enderecoRecebido := TipoDeEndereco(cenario.endInserido)
		if enderecoRecebido != cenario.endEsperado {
			t.Errorf("tipos diferentes, esperava %s e recebeu %s",
				cenario.endEsperado,
				enderecoRecebido)
		}
	}
}

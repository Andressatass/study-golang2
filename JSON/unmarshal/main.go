package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` //a tag define qual o nome da chave vai ser
	Raça  string `json:"raça"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroJSON := `{"nome":"rex","raça":"dalmata","idade":3}`

	var c cachorro

	err := json.Unmarshal([]byte(cachorroJSON), &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

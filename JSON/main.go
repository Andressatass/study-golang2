package main

import (
	"bytes"
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
	c := cachorro{"rex", "dalmata", 3}
	fmt.Println(c)

	cachorroJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroJson)

	fmt.Println(bytes.NewBuffer(cachorroJson))

}

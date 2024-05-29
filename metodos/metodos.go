package main

import "fmt"

type user struct {
	idade int
}

// tem diferença de com e sem ponteiro
func (u *user) fazerAniversario() {
	u.idade++
}

func main() {
	user := user{30}

	user.fazerAniversario()

	fmt.Println(user.idade)
}

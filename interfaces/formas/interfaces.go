package main

import "fmt"

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Println("A area da forma é", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// type circulo stuct {
// 	raio float64
// }

func main() {
	r := retangulo{10, 15}
	escreverArea(r)
}

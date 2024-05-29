package main

import "fmt"

func main() {
	var array1 [5]int
	array1[1] = 1

	fmt.Println(array1)

	array2 := [...]int{1, 2, 3} //tamanho baseado com valores passados
	fmt.Println(array2)

	// SLICE
	slice := []int{1, 2, 3, 4, 5} //n é um array

	slice = append(slice, 18)

	fmt.Print(slice)

	// um slice pode ser um ponteiro para um array
	slice2 := array2[0:1]
	fmt.Println("esse é o slice 2 do array", slice2)

	array2[0] = 4
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("arrays internos ------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(cap(slice3))
}

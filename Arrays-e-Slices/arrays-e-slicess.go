package main

import "fmt"

func main() {

	fmt.Println("arrays e slices")
	var array1 [5]int

	array1[0] = 1
	array1[1] = 2

	fmt.Println(array1)
	//array não é muito utilizado, utilizamos mais o slices pois é mais flexivel no tamanho

	slice := []int{20, 22, 23, 24, 25, 26, 27, 28}

	fmt.Println(slice)

	slice = append(slice, 19) //aqui estamos adicionando mais um numero em uma posição nova
	fmt.Println(slice)
}

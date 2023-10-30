package main

import "fmt"

func main() {

	var valor int8 = 10
	var valor2 int8 = valor

	fmt.Println(valor, valor2)

	//ponteiros é uma referencia de memoria

	var valor3 int

	var ponteiro *int

	valor3 = 100
	ponteiro = &valor3

	fmt.Println(valor3, ponteiro)

	fmt.Println(valor3, *ponteiro) //desreferenciação aqui conseguimos olhar o valor que tem dentro do ponteiro

}

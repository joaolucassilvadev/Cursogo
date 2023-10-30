package main

import "fmt"

//podemos passar varios numeros de uma vez
func sum(numeros ...int) int {

	soma := 0

	for _, numero := range numeros {

		soma += numero
	}

	return soma

}

func main() {

	result := sum(200, 300, 4000, 500, 600, 7000, 80000)

	fmt.Println(result)

}

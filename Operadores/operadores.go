package main

import (
	"fmt"
)

func main() {
	// não podemos usar variaveis de tipos diferentes
	//por exemplo int8 com int32 não podemos utilizar eles juntos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicasao := 10 * 5
	restodivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicasao, restodivisao)

	//operadores de atribuição

	var variavel string = "string"

	variavel2 := "string2"

	fmt.Println(variavel, variavel2)

}

package main

import (
	"fmt"
)

func certo(a, b int) (int, string) {

	if a == b {
		return a + b, "são iguais"

	}

	if a > b {
		return a, "a são maiores"
	}

	if a < b {
		return b, "b é maior que a "
	}

	return 0, "nem uma das opções acima"

}

func somar(n1 int32, n2 int32, n3 int32) int32 {

	return n1 + n2 + n3
}

func calculos(b1, b2 int32) (int32, int32) {

	soma := b1 + b2

	sub := b1 - b2

	return soma, sub
}

func calculo(c2, c1 int32) (int32, int32, int32) {

	soma1 := c2 * (c2 * c1)

	soma2 := c1 * (c2 + c1)

	soma3 := c2 + c1 - c1

	return soma1, soma2, soma3
}

func main() {
	soma := somar(10, 20, 30)

	resultado, mensagem := certo(10, 30)

	fmt.Println(soma)

	calculo1, calculo2 := calculos(30, 50)

	fmt.Println(calculo1, calculo2)
	soma4, soma5, soma6 := calculo(60, 60)
	fmt.Println(soma4, soma5, soma6)

	fmt.Println(resultado, mensagem)

}

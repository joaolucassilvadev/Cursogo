package main

import "fmt"

func n(numero1 int) int {

	soma := 0

	soma += numero1 * 2

	return soma
}

func main() {

	//nesse expemplo sempre que puxarmos a main primeiro já mandamos o valor para n já multiplicado

	somar := func() int {

		return n(10) * 2
	}

	fmt.Println(somar())

}

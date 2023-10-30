package main

import (
	"fmt"
)

func main() {

	var variavel1 string = "variavel 1"

	//tem esse outro jeito de dar valor para uma variavel

	variavel2 := "variavel 2"

	fmt.Println(variavel1)

	fmt.Println(variavel2)

	var (
		variavel3 string = "joao"
		variavel4 string = "lucas"
	)
	fmt.Println(variavel3, variavel4)
	variavel1, variavel2 = variavel2, variavel1

	fmt.Println(variavel1, variavel2)

}

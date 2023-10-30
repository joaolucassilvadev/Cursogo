package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("escrevendo arquivo")
	//para usarmos um pacote sempre pegamos o ultimo nome depois da barra "/"
	auxiliar.Escrever()
	auxiliar.Escrever2()

	erro := checkmail.ValidateFormat("js0454261@gmail.com")

	fmt.Println(erro)
}

package main

import "fmt"

type Endereco struct {
	cidade string
	numero string
	rua    string
}

type cliente struct {
	nome  string
	idade int

	enderecocliente Endereco
}

func main() {

	joao := cliente{
		nome:  "joao",
		idade: 40,
	}

	joao.enderecocliente.cidade = "sao luis"

	fmt.Println(joao)

}

package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

type endereco struct {
	numero int8
	rua    string
}

func main() {
	fmt.Println("arquivo structs")
	var u usuario
	u.nome = "davi"
	u.idade = 21
	fmt.Println(u)

	enderecoex := endereco{32, "rua paraiva"}
	usuario2 := usuario{"davi", 21}

	fmt.Println(usuario2)
	fmt.Println(enderecoex)

}

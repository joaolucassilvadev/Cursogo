package main

import "fmt"

type renda struct {
	rendabru int

	rendali int
}

type Endereco struct {
	cidade string
	numero string
	rua    string
}

type cliente struct {
	nome  string
	idade int

	ativo bool

	enderecocliente Endereco
	salario         renda
}

//aqui estamos ataxado as informações do struct cliente, agora enves de utilizarmos cliente.nome  vamos utilizar c.nome
func (c cliente) inativo() {

	c.ativo = false

	fmt.Printf("o cliente %s foi desativado", c.nome)

}

func main() {

	joao := cliente{
		nome:  "joao",
		idade: 40,
	}

	joao.enderecocliente.cidade = "sao luis"
	joao.salario.rendabru = 1000

	fmt.Println(joao)

	joao.inativo()

}

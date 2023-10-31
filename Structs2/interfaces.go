package main

import "fmt"

type endereco struct {
	cidade string
	numero int
}

type cliente struct {
	nome    string
	salario int
	idade   int
	endereco
	ativo bool
}

func (c cliente) Desativar() {
	c.ativo = false

	fmt.Printf("o cliente %s foi desativado", c.nome)
}

//a interface permiti uma abistração maior das structs, por exemplo sempre podemos utilizar o metodo desativar se a struct tiver
// uma func desativar, assim deixa o codigo mais simples como visto abaixo
//toda interface só recebe metodos
type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {

	pessoa.Desativar()
}

func main() {

	wesley := cliente{
		nome:    "joao",
		idade:   10,
		salario: 1000,
	}
	fmt.Println(wesley)

	Desativacao(wesley)

}

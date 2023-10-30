package main

import "fmt"

//aprendemos sobre struct encadeadas
type Endereco struct {
	cep        int
	endereco   string
	numero     string
	cidade     string
	estado     string
	logradouro string
}

// se eu tenho duas structs e preciso que uma informação
//de uma esteja dentro de outra, fazemso da sequinte forma
type cliente struct {
	nome  string
	idade int

	Endereco //dessa forma as informções da struct Endereco vai tá
	//dentro da struct cliente

	//podemos utilizar a struct como um tipo também

	enderecocliente Endereco
	// assim podemos botar andrees.algo que tenha dentro da struct endereço
	// assim é bom que podemos botar enderecocliente Endereco
}

func main() {

	joao := cliente{
		nome:  "joao",
		idade: 30,
	}
	joao.Endereco.logradouro = "rua paraiba "

	joao.enderecocliente.estado = "maranhão"
	fmt.Println(joao)
}

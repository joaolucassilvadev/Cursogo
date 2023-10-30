package main

//basico basico de structs

import "fmt"

type cliente struct {
	nome string

	idade int

	ativo bool
}

//structs é muito importante vamos utiliazar bastante
func main() {
	wesley := cliente{
		nome:  "wesley",
		idade: 39,
		ativo: true,
	}

	fmt.Println(wesley)
}

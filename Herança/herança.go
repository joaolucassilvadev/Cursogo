package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     string
	altura    string
}

type estudante struct {
	pessoa

	curso     string
	faculdade string
}

func main() {

	fmt.Println("herança")

	p1 := pessoa{"joao", "pedro", "20", "179"}

	fmt.Println(p1)
	e1 := estudante{p1, "engenharia", "ups"}

	fmt.Println(e1)

	fmt.Println(e1.altura)

}

package main

import "fmt"

func main() {

	fmt.Println("maps em go")

	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "silva",
	}

	salario := map[string]int{
		"nome": 1000,
		"rafa": 2000,
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(salario["rafa"])

	for nome, salario := range salario {
		fmt.Printf("o salario de %s é de %d\n", nome, salario)
		//println não aparece o nome e nem salario, só aparece a mensagem toda
	}

	for _, salario := range salario {
		fmt.Printf("o salario de é de %d\n", salario)
		//se utilizarmos um _ não precisamos botar a variavel nome se a gente quiser exibir somente uma variavel

	}
}

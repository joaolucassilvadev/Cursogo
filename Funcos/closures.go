package main

//quando temos funções anonimas conseguimos rodar funcções dentro de outras funções
import "fmt"

func soma(valor ...int) int {

	d := 0
	for _, soma := range valor {

		d += soma

	}

	return d

}

func main() {

	d := func() int {
		return soma(10, 20, 30, 40) * 2

		// por exemplo nesse exemplo estamos mandando o valor
		//para a variavel soma voltando multiplicando por 2 e imprimindo o resultado na main
	}
	fmt.Println(d())
	valores := soma(10, 20, 30, 40, 506)

	fmt.Println(valores)

}

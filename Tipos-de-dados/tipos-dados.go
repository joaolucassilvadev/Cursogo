package make

import (
	"errors"
	"fmt"
)

func main() {

	//int8, int16, int32 , int64

	// se botarmos somente int ele vai utilizar a arquitetura do meu computador
	//se meu computador for 64 bits ele vai usar 64 se ele for 32 ele vai usar 32 bits

	var numero int8 = 100

	numero2 := 1000

	fmt.Println(numero)

	fmt.Println(numero2)

	//numeros reais

	//float32,float64
	//não podemos declarar somente float
	var numero3 float64 = 100.00

	fmt.Println(numero3)

	//strings

	var nome string = "joao"

	fmt.Println(nome)

	var erro error = errors.New("Erro interno")

	fmt.Println(erro)

}

package main

import "fmt"

func main() {
	a := 10

	fmt.Println(&a)        //aqui achamos o endereco a qual esta armazenado meu numero
	fmt.Println(*&a)       //aqui apontamos para o ponteiro para saber o que ele armazena
	var ponteiro *int = &a //aqui fazemos uma variavel ponteiro do tipo int com o valor do endereço de a

	*ponteiro = 20

	fmt.Println(ponteiro) //podemos observar quando atribuimos ponteiro ao valor de a  atribui o endereço de a também
	fmt.Println(*&a)      //já que atribuimos o endereço do ponteiro a na variavel ponteiro e botamos *ponteiro =20 logo no endereço o valor 10 é trocado
	//pelo valor 20 ou seja sempre que apontarmos para o endereço a de novo e pedimos o valor de lá vamos receber o valor 20 apartir de agora
	b := &a

	fmt.Println(b) //podemos observar quando atribuimos b ao valor de a b atribui o endereço de a também
}

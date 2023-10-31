package main

import "fmt"

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}
func soma1(a, b int) int {

	a = 50

	return a + b
}

func main() {
	variavel1 := 10
	variavel2 := 15

	variavel3 := 20
	variavel4 := 30

	soma(&variavel1, &variavel2)
	soma1(variavel3, variavel4)
	fmt.Println(variavel1)
	fmt.Print(variavel3)

}

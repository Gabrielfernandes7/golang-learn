package main

import "fmt"

func main()  {
	x := 10
	var ponteiro *int
	ponteiro = &x
	fmt.Println("Valor de x: ", x)
	fmt.Println("Endereço de x: ", &x)
	fmt.Println("Valor de ponteiro: ", ponteiro)
	fmt.Println("Valor apontado por ponteiro: ", *ponteiro)
}
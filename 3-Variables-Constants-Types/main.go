package main

import "fmt"

func main()  {
	fmt.Println("Declaração de variáveis")

	var a, b int = 20, 10
	var soma int

	const PI float32 = 3.1415
	var perimetro float32 = PI*2*float32(a);

	soma = a + b

	fmt.Println(soma)
	fmt.Print(perimetro)
}
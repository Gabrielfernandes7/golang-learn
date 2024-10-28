package main

import "fmt"

func main()  {
	fmt.Println("Conversão de base")

	numero_base_10 := 2548

	// prefix 0x represent hexadecimal
	numero_base_16 := 0x9F4

	// lower case output
	fmt.Printf("%x \n", numero_base_10)

	// uppercase output
	fmt.Printf("%X \n", numero_base_10)

	fmt.Printf("%d \n", numero_base_16)
	fmt.Printf("%o \n", numero_base_10)

}
package main

import "fmt"

func main() {
	const emailToSend int = 3
	emailSent := 0

	for emailSent < emailToSend {
		fmt.Println("Email sendo enviado...", emailSent)
		emailSent++
	}

	fmt.Println("Fim do sistema")
}
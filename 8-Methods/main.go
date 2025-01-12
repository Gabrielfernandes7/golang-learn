package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
}

func (p *Pessoa) saudar() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos. \n", p.Nome, p.Idade)
}

func main() {
	pessoa := Pessoa {
		Nome: "João",
		Idade: 50,
	}

	pessoa.saudar()
}
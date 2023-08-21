package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main() {
	var p Pessoa
	p = Pessoa{
		Nome: "Carol",
	}

	var pessoas [3]Pessoa

	pessoas[0] = p
	for _, pessoa := range pessoas { // _ é para desconsiderar o índice e range é para passar por cada elemento dentro do array
		if (pessoa != Pessoa{}) { // Quais elementos no array não estão
			fmt.Println(pessoa)

		}
	}

	pessoas = adiciona("Pedro", pessoas)
	fmt.Println(pessoas)
}

func adiciona(nome string, a [3]Pessoa) [3]Pessoa {
	a[0] = Pessoa{Nome: nome}
	return a
}

package main

import "fmt"

type Carro struct {
	Modelo     string
	Velocidade int
	Cor        string
	Ligado     bool
}

func (c *Carro) Ligar() {
	c.Ligado = !c.Ligado
}

func (c *Carro) Acelerar(valor int) {
	c.Velocidade += valor
}

func main() {
	carro := Carro{
		Modelo:     "Defender",
		Velocidade: 180,
		Cor:        "Preto fosco",
		Ligado:     true,
	}

	fmt.Println(carro)
	carro.Ligar()
	fmt.Println(carro)
	carro.Acelerar(10)
	fmt.Println(carro)
}

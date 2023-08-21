package main

import "fmt"

type Carro struct {
	Modelo     string
	Velocidade int
	Cor        string
	EstaLigado bool
}

func (c *Carro) Ligar() {
	c.EstaLigado = true
}

func (c *Carro) Desligar() {
	c.EstaLigado = true
	c.Velocidade = 0
}

func (c *Carro) Acelerar(valor int) {
	c.Velocidade += valor
}

func main() {
	carro := Carro{
		Modelo:     "HB20",
		Velocidade: 180,
		Cor:        "Cinza",
		EstaLigado: true,
	}

	carro.Ligar()
	fmt.Println(carro)

	carro.Acelerar(50)
	fmt.Println(carro)

	carro.Desligar()
	fmt.Println(carro)
}

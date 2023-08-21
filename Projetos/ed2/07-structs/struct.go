package main

import (
	"fmt"
	"math"
)

type Pessoa struct {
	Nome   string
	Idade  int
	Altura float64
}

func (p *Pessoa) AvancaIdade() {
	p.Idade++
}

type Ponto struct {
	X, Y int
}

type Retangulo struct {
	Ponto
	Largura, Altura int
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {
	p := Pessoa{
		Nome:   "Alice",
		Idade:  25,
		Altura: 1.65,
	}

	fmt.Println(p)
	fmt.Println(p.Nome)
	fmt.Println(p.Altura)

	r := Retangulo{
		Ponto:   Ponto{X: 1, Y: 2},
		Largura: 60,
		Altura:  30,
	}
	fmt.Println(r)

	circ := Circulo{
		Raio: 1.5,
	}
	fmt.Println(circ.Area())
}

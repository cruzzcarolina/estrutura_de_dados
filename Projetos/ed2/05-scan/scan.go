package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x, z int
	var y float64

	fmt.Println("Por favor informe dois valores inteiros: ")
	fmt.Scanln(&x, &z)
	fmt.Println(x)
	fmt.Println(z)

	fmt.Println("Informe um float")
	fmt.Scanln(&y)
	fmt.Println(y)

	// Lendo frases inteiras
	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("Escreva um texto: ")
	msg, _ := leitor.ReadString('\n')

	fmt.Println(msg)
}

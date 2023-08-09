package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um valor inteiro de 1 a 7: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Erro ao ler o número:", err)
		return
	}

	dia := diaEscrito(n)
	fmt.Println(dia)
}

func diaEscrito(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Valor inválido!"
	}
}

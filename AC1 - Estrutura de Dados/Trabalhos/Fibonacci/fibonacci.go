package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite a um número que representa a sequência em Fibonacci: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Erro ao ler o número:", err)
		return
	}

	result := fibo(n)
	fmt.Printf("O %d-ésimo elemento da série de Fibonacci é: %d\n", n, result)
}

func fibo(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	ant1 := 1
	ant2 := 1

	for i := 3; i <= n; i++ {
		current := ant1 + ant2
		ant2 = ant1
		ant1 = current
	}

	return ant1
}

package main

import "fmt"

func main() {
	var qtdeDiscos int
	fmt.Print("escreva o número de discos: ")
	fmt.Scan(&qtdeDiscos)

	fmt.Println()
	totalMovimentos := hanoi(qtdeDiscos, "A", "C", "B")

	fmt.Printf("resolvido em %d movimentos", totalMovimentos)
}

func hanoi(n int, origem, destino, apoio string) int {
	if n == 1 {
		fmt.Printf("move peça 1: %s -> %s \n", origem, destino)
		return 1
	}
	movimentos := hanoi(n-1, origem, apoio, destino)
	fmt.Printf("move peça %d: %s -> %s \n", n, origem, destino)
	movimentos++
	movimentos += hanoi(n-1, apoio, destino, origem)
	return movimentos
}

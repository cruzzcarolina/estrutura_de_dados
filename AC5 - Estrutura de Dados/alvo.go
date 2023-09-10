package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	alvo := 10
	esquerda, direita := encontrarsoma(array, alvo)
	if esquerda == -1 && direita == -1 {
		fmt.Println("nenhum par encontrado")
	} else {
		fmt.Printf("par encontrado: (%d, %d)", esquerda, direita)
	}
}

func encontrarsoma(array []int, alvo int) (int, int) {
	esq := 0
	dir := len(array) - 1

	for esq < dir {
		soma := array[esq] + array[dir]
		if soma == alvo {
			return array[esq], array[dir]
		} else if soma < alvo {
			esq++
		} else {
			dir--
		}
	}

	return -1, -1
}

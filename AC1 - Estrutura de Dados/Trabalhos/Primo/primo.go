package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	e_primo(n)
}

func e_primo(n int) {
	e_primo := false
	divisores := []int{}

	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			e_primo = true
			divisores = append(divisores, i)
		}
	}

	if e_primo {
		fmt.Printf("%d não é primo. Divisores: %v\n", n, divisores)
	} else {
		fmt.Printf("%d é primo.\n", n)
	}
}

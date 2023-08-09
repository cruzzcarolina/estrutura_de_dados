package main

import "fmt"

func e_bissexto(a int) bool {
	if a%4 == 0 {
		if a%100 != 0 || a%400 == 0 {
			return true
		}
	}
	return false
}

func main() {
	var a int
	fmt.Print("Escreva um ano: ")
	fmt.Scan(&a)

	if e_bissexto(a) {
		fmt.Printf("%d é um ano bissexto.\n", a)
	} else {
		fmt.Printf("%d não é um ano bissexto.\n", a)
	}
}

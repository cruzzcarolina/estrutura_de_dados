package main

import "fmt"

func main() {
	var numero int
	maior := 0

	for {
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}
		if numero > maior {
			maior = numero
		}
	}

	fmt.Println(maior)
}

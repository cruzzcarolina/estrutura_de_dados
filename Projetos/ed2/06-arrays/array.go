package main

import "fmt"

func main() {
	// Array é uma coleção de dados do mesmo tipo
	// Possui tamanho fixo

	var filmes [5]string

	filmes[0] = " O Senhor dos Anéis"
	filmes[1] = "O Poderoso Chefão"
	filmes[2] = "Barbie"
	fmt.Println(filmes)

	// É possível usarmos declaração curta
	numeros := [4]int{0, 2, 4, 6}
	fmt.Println(numeros)

	// Slices são estruturtas mais flexíveis
	// Possuem tamanho dinâmico

	var novosNumeros []int
	novosNumeros = numeros[1:]  // Vai até o final do array
	novosNumeros = numeros[1:3] // Não inclui o elemento 3
	fmt.Println(novosNumeros)

	fmt.Printf("%p", &numeros) // %p é para indicar o endereço da memória
	fmt.Println("")
	fmt.Printf("%p", &novosNumeros)

	// É possível usar declarações curtas
	nomes := []string{"Ana", "Pedro"}
	fmt.Println(nomes)

	// Estruturas de repetição com arrays ou slices
	fmt.Println("------------------")
	for i := 0; i < len(numeros); i++ { // len é o tamanho do array
		fmt.Println(numeros[i])
	}

	// range é similar ao enumerate() em Python
	for indice, num := range numeros {
		fmt.Println("índice: ", indice, "- valor", num)
	}

	fmt.Println("-----------------")
	fmt.Println(numeros)
	modificarArray(numeros) // Não altera o array numeros!
	fmt.Println(numeros)

	fmt.Println(novosNumeros)
	modificarSlice(novosNumeros) // Altera o slice original! Slice é fatia do array.
	fmt.Println(novosNumeros)

	fmt.Println(numeros)

	maisNumeros := criarSlice()
	fmt.Println(maisNumeros)

}

// Em Go, quando passamos um array como parâmetro, criamos uma cópia deste array.
func modificarArray(a [4]int) {
	a[0] = 999
}

// Em slices, o original também é alterado.
func modificarSlice(s []int) {
	s[0] = 999
}

func criarSlice() []int {
	novoSlice := []int{0, 20, 30}
	return novoSlice
}

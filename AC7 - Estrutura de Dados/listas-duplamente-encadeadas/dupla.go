package main

import "fmt"

type No struct {
	value    int
	proximo     *No
	anterior *No
}

type ListaDupla struct {
	primeiro *No
}

func main() {
	lista := &ListaDupla{}
	inserirNo(lista, 1, 1)
	inserirNo(lista, 2, 2)
	inserirNo(lista, 3, 3)
	exibirLista(lista)

	fmt.Println("removendo o nó com valor 2:")
	removerNo(lista, 2)
	exibirLista(lista)
}

func exibirLista(lista *ListaDupla) {
	noAtual := lista.primeiro
	for noAtual != nil {
		fmt.Println(noAtual.value)
		noAtual = noAtual.proximo
	}
}

func inserirNo(lista *ListaDupla, valor, posicao int) {
	novoNo := &No{value: valor}
	if lista.primeiro == nil {
		lista.primeiro = novoNo
	} else {
		noAtual := lista.primeiro
		for i := 1; i < posicao-1; i++ {
			if noAtual == nil {
				fmt.Println("posição inválida")
				return
			}
			noAtual = noAtual.proximo
		}
		novoNo.proximo = noAtual.proximo
		novoNo.anterior = noAtual
		if noAtual.proximo != nil {
			noAtual.proximo.anterior = novoNo
		}
		noAtual.proximo = novoNo
	}
}

func removerNo(lista *ListaDupla, valor int) {
	noAtual := lista.primeiro
	for noAtual != nil {
		if noAtual.value == valor {
			if noAtual.anterior != nil {
				noAtual.anterior.proximo = noAtual.proximo
			} else {
				lista.primeiro = noAtual.proximo
			}
			if noAtual.proximo != nil {
				noAtual.proximo.anterior = noAtual.anterior
			}
			return
		}
		noAtual = noAtual.proximo
	}
	fmt.Println("nó não encontrado")
}

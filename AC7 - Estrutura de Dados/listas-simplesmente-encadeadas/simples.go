package main

import "fmt"

type No struct {
	valor   int
	proximo *No
}

type Lista struct {
	primeiro *No
}

func main() {
	lista := &Lista{}
	inserirNoInicio(lista, 1)
	inserirNoInicio(lista, 2)
	inserirNoInicio(lista, 3)
	exibirListaCircular(lista)

	fmt.Println("excluindo o primeiro nó:")
	excluirPrimeiroNo(lista)
	exibirListaCircular(lista)
}

func exibirListaCircular(lista *Lista) {
	if lista.primeiro == nil {
		fmt.Println("lista vazia")
		return
	}
	noAtual := lista.primeiro
	for {
		fmt.Println(noAtual.valor)
		noAtual = noAtual.proximo
		if noAtual == lista.primeiro {
			break
		}
	}
}

func inserirNoInicio(lista *Lista, valor int) {
	novoNo := &No{valor: valor}
	if lista.primeiro == nil {
		lista.primeiro = novoNo
		novoNo.proximo = lista.primeiro
	} else {
		novoNo.proximo = lista.primeiro
		noAtual := lista.primeiro
		for noAtual.proximo != lista.primeiro {
			noAtual = noAtual.proximo
		}
		noAtual.proximo = novoNo
		lista.primeiro = novoNo
	}
}

func excluirPrimeiroNo(lista *Lista) {
	if lista.primeiro == nil {
		fmt.Println("lista vazia")
		return
	}

	if lista.primeiro.proximo == lista.primeiro {
		lista.primeiro = nil
	} else {
		noAtual := lista.primeiro
		for noAtual.proximo != lista.primeiro {
			noAtual = noAtual.proximo
		}
		noAtual.proximo = lista.primeiro.proximo
		lista.primeiro = lista.primeiro.proximo
	}
}

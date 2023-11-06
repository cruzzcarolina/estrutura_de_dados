package main

import "fmt"

type Node struct {
	Chave    int
	Esquerda *Node
	Direita  *Node
}

type Arvore struct {
	Raiz *Node
}

func main() {
	arvore := Arvore{}
	arvore.inserir(50)
	arvore.inserir(30)
	arvore.inserir(70)
	arvore.inserir(20)
	arvore.inserir(40)
	arvore.inserir(60)
	arvore.inserir(80)
	fmt.Println(arvore.buscar(30))
	fmt.Println(arvore.buscar(45))
}

func (a *Arvore) inserir(chave int) {
	if a.Raiz == nil {
		a.Raiz = &Node{Chave: chave}
	} else {
		a.inserindo(a.Raiz, chave)
	}
}

func (a *Arvore) inserindo(no *Node, chave int) {
	if chave < no.Chave {
		if no.Esquerda == nil {
			no.Esquerda = &Node{Chave: chave}
		} else {
			a.inserindo(no.Esquerda, chave)
		}
	} else if chave > no.Chave {
		if no.Direita == nil {
			no.Direita = &Node{Chave: chave}
		} else {
			a.inserindo(no.Direita, chave)
		}
	}
}

func (a *Arvore) buscar(chave int) bool {
	return a.buscando(a.Raiz, chave)
}

func (a *Arvore) buscando(no *Node, chave int) bool {
	if no == nil {
		return false
	}
	if chave == no.Chave {
		return true
	} else if chave < no.Chave {
		return a.buscando(no.Esquerda, chave)
	} else {
		return a.buscando(no.Direita, chave)
	}
}

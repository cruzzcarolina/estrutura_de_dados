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
	arvore.Inserir(50)
	arvore.Inserir(30)
	arvore.Inserir(70)
	arvore.Inserir(20)
	arvore.Inserir(40)
	arvore.Inserir(60)
	arvore.Inserir(80)

	fmt.Println(arvore.Buscar(30))
	fmt.Println(arvore.Buscar(45))
}

func (a *Arvore) Inserir(chave int) {
	if a.Raiz == nil {
		a.Raiz = &Node{Chave: chave}
	} else {
		a.inserir(a.Raiz, chave)
	}
}

func (a *Arvore) inserir(no *Node, chave int) {
	if chave < no.Chave {
		if no.Esquerda == nil {
			no.Esquerda = &Node{Chave: chave}
		} else {
			a.inserir(no.Esquerda, chave)
		}
	} else if chave > no.Chave {
		if no.Direita == nil {
			no.Direita = &Node{Chave: chave}
		} else {
			a.inserir(no.Direita, chave)
		}
	}
}

func (a *Arvore) Buscar(chave int) bool {
	return a.buscarChaveRec(a.Raiz, chave)
}

func (a *Arvore) buscarChaveRec(no *Node, chave int) bool {
	if no == nil {
		return false
	}
	if chave == no.Chave {
		return true
	} else if chave < no.Chave {
		return a.buscarChaveRec(no.Esquerda, chave)
	} else {
		return a.buscarChaveRec(no.Direita, chave)
	}
}

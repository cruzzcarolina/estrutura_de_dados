package operacao

import "fmt"

func AdicionaContato(c *contato.Contato, a *[5]contato.Contato) {
	for ind, cont := range a {
		if cont == (contato.Contato{}) {
			a[ind] = *c
			break
		}
	}
}

func ExcluiContato(a *[5]contato.Contato) {
	if a[0] == (contato.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
		return
	}

	for ind, cont := range a {
		if cont == (contato.Contato{}) {
			a[ind-1] = contato.Contato{}
		}
	}
}

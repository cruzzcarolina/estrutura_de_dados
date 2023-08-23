package main

import "fmt"

func main() {
	var contatos [5]contato.Contato
	var nome, email, opcao string
	var indice int

	leitor := NewReader(Stdin)

	fmt.Println("lista de contatos: ")
	for {
		fmt.Print("Escreva 1 para adicionar, 2 para remover, 3 para editar email, 4 para mostrar contatos ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("escreva seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("escreva seu email: ")
			fmt.Scanln(&email)

			operacao.AdicionaContato(&contato.Contato{Nome: nome, Email: email}, &contatos)
		case "2":
			operacao.ExcluiContato(&contatos)
		case "3":
			fmt.Print("índice do contato a ser editado: ")
			fmt.Scanln(&indice)

			fmt.Print("fale o novo email: ")
			fmt.Scanln(&email)

			operacao.EditaEmail(indice, email, &contatos)
		case "4":
			fmt.Println("contatos:")
			for i, c := range contatos {
				if c != (contato.Contato{}) {
					fmt.Printf("Nome: %s, Email: %s\n", i, c.Nome, c.Email)
				}
			}
		default:
			return
		}
	}
}

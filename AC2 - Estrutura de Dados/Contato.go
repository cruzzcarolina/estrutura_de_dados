package main

import "fmt"

// Elabore um struct Contato, que deve conter os campos nome e e-mail. O struct deve conter um método para alterar o e-mail.
// Elabore uma função adicionaContato, que recebe um contato e um array com até 5 elementos e inclui o contato no primeiro elemento vazio do array.
// Elabore uma função excluiContato, que recebe um array de 5 elementos e elimina o último contato válido.
// Elabore uma interface por linha de comando na função main, que cria um array de 5 elementos e permite a inclusão ou exclusão de contatos.

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(email string) {
	c.Email = email
}

func adicionaContato(contato Contato, array []Contato) []Contato {
	for i := range array {
		if array[i].Nome == "" {
			array[i] = contato
			break
		}
	}
	return array
}

func excluiContato(array []Contato) []Contato {
	for i := len(array) - 1; i >= 0; i-- {
		if array[i].Nome != "" {
			array[i] = Contato{}
			break
		}
	}
	return array
}

func main() {
	contato := Contato{
		Nome:  "Carol",
		Email: "carol.cruz2708@gmail.com",
	}

	contato.AlterarEmail("carolzinha@gmail.com")
	fmt.Println(contato)

	arrayContatos := make([]Contato, 5)

	var nome, email string
	fmt.Print("nome: ")
	fmt.Scanln(&nome)
	fmt.Print("email: ")
	fmt.Scanln(&email)

	novoContato := Contato{Nome: nome, Email: email}
	arrayContatos = adicionaContato(novoContato, arrayContatos)
}

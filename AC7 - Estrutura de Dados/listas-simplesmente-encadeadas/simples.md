# Lista Circular Simplesmente Encadeada

Uma lista circular simplesmente encadeada é uma estrutura de dados na qual os nós estão conectados de forma circular, ou seja, o último nó aponta para o primeiro nó. Os pseudocódigos:

## Exibição dos nós em uma lista circular

```
exibirListaCircular(lista):
    se lista estiver vazia:
        retorne "lista vazia"
    no_atual = lista.primeiro
    faça:
        exiba no_atual.valor
        no_atual = no_atual.proximo
    enquanto no_atual for diferente de lista.primeiro
```
## Inserção de um nó no início da lista
```
inserirNoInicio(lista, valor):
    novo_no = criar_nó(valor)
    se lista estiver vazia:
        lista.primeiro = novo_no
        lista.primeiro.proximo = lista.primeiro
    senão:
        novo_no.proximo = lista.primeiro
        no_atual = lista.primeiro
        enquanto no_atual.proximo for diferente de lista.primeiro:
            no_atual = no_atual.proximo
        no_atual.proximo = novo_no
        lista.primeiro = novo_no
```
## Exclusão do primeiro nó da lista
```
inserirNoInicio(lista, valor):
    novo_no = criar_nó(valor)
    se lista estiver vazia:
        lista.primeiro = novo_no
        lista.primeiro.proximo = lista.primeiro
    senão:
        novo_no.proximo = lista.primeiro
        no_atual = lista.primeiro
        enquanto no_atual.proximo for diferente de lista.primeiro:
            no_atual = no_atual.proximo
        no_atual.proximo = novo_no
        lista.primeiro = novo_no
```
# Lista Duplamente Encadeada

Uma lista duplamente encadeada é uma estrutura de dados em que cada nó tem referências tanto para o nó anterior quanto para o nó seguinte. Os pseudocódigos:

## Exibição dos nós em uma lista duplamente encadeada
```
exibirListaDuplamenteEncadeada(lista):
    no_atual = lista.primeiro
    enquanto no_atual for diferente de nulo:
        exiba no_atual.valor
        no_atual = no_atual.próximo
```
## Busca de um nó
```
buscarNodo(lista, valor):
    no_atual = lista.primeiro
    enquanto no_atual for diferente de nulo e no_atual.valor for diferente valor:
        no_atual = no_atual.próximo
    se no_atual for diferente de nulo:
        retorne no_atual
    senão:
        retorne nulo
```
## Inserção de um nó
```
inserirNo(lista, valor, posição):
    novo_no = criar_no(valor)
    no_atual = lista.primeiro
    para i de 1 até posição-1:
        se no_atual for diferente nulo:
            no_atual = no_atual.proximo
        senão:
            retorne "Posição inválida"
    novo_no.anterior = no_atual
    novo_no.proximo = no_atual.proximo
    se no_atual.proximo for diferente nulo:
        no_atual.proximo.anterior = novo_no
    no_atual.proximo = novo_no
```
## Remoção de um nó
```
removerNo(lista, valor):
    no_alvo = buscarNo(lista, valor)
    se no_alvo for diferente de nulo:
        no_anterior = no_alvo.anterior
        no_seguinte = no_alvo.proximo
        se no_anterior for diferente nulo:
            no_anterior.proximo = no_seguinte
        senão:
            lista.primeiro = no_seguinte
        se no_seguinte for diferente nulo:
            no_seguinte.anterior = no_anterior
    senão:
        retorne "nó não encontrado"
```
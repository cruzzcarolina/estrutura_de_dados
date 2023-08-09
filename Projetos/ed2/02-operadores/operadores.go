package main

import "fmt"

func main(){
	var x,y, z = 4, 10, -2

	//aritmeticos
	fmt.println(x + y)
	fmt.println(x - y)
	fmt.println(x * y)
	fmt.println(x / y)
	fmt.println(x % y)

	//atribuicao
	z++
	z--
	z *= 2
	z += 2
	z -= 2
	z /= 2
	z %= 2

	//comparacao
	fmt.println(x == y)
	fmt.println(x != y)
	fmt.println(x > y)
	fmt.println(x >= y)
	fmt.println(x < y)
	fmt.println(x <= y)

	//logicos
	fmt.println(true && false) //AND
	fmt.println(true || false) //OR
	fmt.println(!true) //NOT

	// não vamos falar
	//bitwise
	//memoria e canal




}


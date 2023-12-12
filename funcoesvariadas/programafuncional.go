package main

import "fmt"

func main(){

	resultado := func(x ...int) func() int{
		res := 0

		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}

	// Aqui retornamos uma função portanto ela emitira no console um ponto na memória
	// fmt.Println(resultado (54,54,54,54))

	// Para retornar o valor em si é necessario usar () no fim da função ficando como abaixo
	fmt.Println(resultado (54,54,54,54)())
	
}

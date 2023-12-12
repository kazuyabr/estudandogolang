package main

import (
	"fmt"
)

func main(){

	variavel := 10
	fmt.Println("Valor da variavel antes da função abc: ", variavel)
	//usando agora uma função para alterar valores na memória
	abc(&variavel)

	fmt.Println("Valor da variavel depois da função abc: ", variavel)
}

// esta função aponta para um endereço de memória e o altera seu valor armazenado
func abc(a *int){
	*a = 200
}
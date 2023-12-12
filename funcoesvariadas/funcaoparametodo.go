package main

import "fmt"

// Jeito GO de ser usa-se de estruturas para representação de objetos
type Carro struct {
	Nome string
}

// usa-se um attach ou bind para referenciar o metodo que criamos usando-se de variavel e tipo como normalmente outras linguagens fazem como o java
// Pode-se dizer que a função abaixo é um metodo responsavel por fazer o nosso objeto carro andar.
func(c Carro) andar(){
	fmt.Println(c.Nome, "andou")
}

func main(){

	carro := Carro {
		Nome: "BMW",
	}
	
	carro.andar()
}

package main

import (
	"fmt"
)

type Carro struct{
	Name string
}

// Usa-se ponteiro para conseguir mudar o valor usando a função.
func (c *Carro) andarPonteiro(){
	c.Name = "BMW"
	fmt.Println(c.Name + " andou!")
}

// Aqui não sera possivel alterar o valor do name pq não esta apontando para a memória
func (c Carro) andarNormal(){
	c.Name = "BMW"
	fmt.Println(c.Name + " andou!")
}

func main(){

	carro := Carro{
		Name: "Fusca",
	}
	//Caso de uso sem apontar para a memória
	carro.andarNormal()
	fmt.Println(carro.Name)

	//caso de uso apontando para a memória
	carro.andarPonteiro()
	fmt.Println(carro.Name)

	fmt.Println("Não tem certo ou errado para o uso de ponteiros aqui é necessario entender apenas os casos de uso.")
}

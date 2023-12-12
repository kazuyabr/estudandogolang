package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome string
	Email string
	CPF string
}

// Supondo que você repita uma estrutura e nao a componha com outra existente
type ClienteNaoComposto struct {
	Nome string
	Email string
	CPF string
	Pais string
}

//Jeito certo aproveitando dados de estrutura ja existente (pode ser entendido como herança essa composição)
type ClienteComposto struct {
	Cliente
	Pais string
}

func (c Cliente) Exibe(){
	fmt.Println("Usando metodo para exibir o cliente: ", c.Nome)
}

func main(){

	cliente := Cliente{
		Nome: "ABC",
		Email: "abc@email.com",
		CPF: "000.000.000-00",
	}

	fmt.Printf("Cliente NORMAL -> Nome: %s. Email: %s. CPF: %s\n", cliente.Nome, cliente.Email, cliente.CPF)

	clienteNaoComposto := ClienteNaoComposto{
		Nome: "DEF",
		Email: "DEF@email.com",
		CPF: "111.111.111-11",
		Pais: "USA",
	}

	fmt.Printf("Cliente Não Composto -> Nome: %s. Email: %s. CPF: %s. Pais: %s\n", clienteNaoComposto.Nome, clienteNaoComposto.Email, clienteNaoComposto.CPF, clienteNaoComposto.Pais)

	clienteComposto := ClienteComposto{
		Cliente: Cliente{
			Nome: "GHI",
			Email: "GHI@email.com",
			CPF: "222.222.222-22",
		},
		Pais: "USA",
	}

	// Entende-se que uma estrtura composta trata sua composição como parte de si mesma uma vez não sendo necessario uso de clienteComposto.Cliente.Nome para acessar nome
	// Acessando diretamente como se houve-se uma desestruturação do objeto ficando em clietneComposto.Nome.
	fmt.Printf("Cliente Composto -> Nome: %s. Email: %s. CPF: %s. Pais: %s\n", clienteComposto.Nome, clienteComposto.Email, clienteComposto.CPF, clienteComposto.Pais)

	// Usando metodo para exibir cliente onde so metodos compostos poderam usar sem ter bind
	cliente.Exibe()
	clienteComposto.Exibe()

	// No caso clienteNaoComposto não foi composto de cliente sendo assim ele precisaria de uma função própria para funcionar
	// clienteNaoComposto.Exibe() // como isso iria falhar ele esta comentado.

	// converter struct para json
	clienteJson, err := json.Marshal(clienteComposto)
	// Tratando o erro
	if err != nil {
		log.Fatal(err.Error())
	}

	// Exibindo o json Marshal tras um slight de byte
	fmt.Println(clienteJson)

	// Ja que ele tras os bytes então temos que converter para string
	fmt.Println(string(clienteJson))

}
package main

import "fmt"

type Cliente struct {
	Nome string
	Email string
	CPF string
}

func main(){

	cliente := Cliente{
		Nome: "ABC",
		Email: "abc@email.com",
		CPF: "000.000.000-00",
	}

	fmt.Println(cliente)
}
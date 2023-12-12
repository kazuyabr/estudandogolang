package main

import "fmt"

type ClienteName string
type ClienteEmail string
type ClientCPF string

func main(){
	ClienteNome := "ABC"
	ClienteEmail := "abc@gmail.com"
	ClienteCPF := "000.000.000-00"

	fmt.Printf("O nome de nosso cliente é: %v\n", ClienteNome)
	fmt.Printf("O email de nosso cliente é: %v\n", ClienteEmail)
	fmt.Printf("O cpf de nosso cliente é: %v\n", ClienteCPF)
}
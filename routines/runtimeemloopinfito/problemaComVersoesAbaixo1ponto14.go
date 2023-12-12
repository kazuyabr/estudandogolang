package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	//vamos garantir q só teremos 1 thread aberta para trabalharmos nesse conceito
	runtime.GOMAXPROCS(1)

	fmt.Println("Começou!")

	//Aqui temos um loop infinito
	//versões anteriores a 1.14 não conseguem reconhecer quando um processo esta sendo usado de modo ocioso, ou seja loop infinito que não faz nada
	//versões 1.13 para tras ficariam travados neste processo e nunca fechariam ou seja nunca seria exibido a frase "Terminou"
	go func(){
		for{

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Terminou!")
}

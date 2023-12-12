package main

import (
	"fmt"
)

func main(){


	// memoria-0xc00000a0c8 <------------- a < ------------- 10 na ocasição que foi rodada antes isso pode mudar

	// Armazenamos 10 a memória
	a := 10
	// Mostramos o endereço da memória
	fmt.Println(&a)

	// Atribuimos um ponteiro para esse endereço de memória usando *int (que pega o valor deste endereço de memória) e &a que pega o endereço da memória
	var ponteiro *int = &a
	*ponteiro = 50

	fmt.Println("Valor de a depois que foi alterada pela variavel ponteiro: ", a)

	//agora alteramos usando a variavel b
	b := &a
	*b = 60
	fmt.Println("valor de a depois que foi alterada pela variavel b", a)
	fmt.Println("note que uma vez q vê aponta para um endereço na memória o valor real sera atualizado sempre enquanto as demais serão só apontamentos (pode ser entendidos como atalhos para alterar um valor pré armazenado!)")


}
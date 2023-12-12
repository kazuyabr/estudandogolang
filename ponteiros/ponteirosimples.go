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
	// Mostra o valor do ponteiro, ou seja, mostra o endereço de memória armazenado no ponteiro usando * antes da variavel ficando *ponteiro
	// se mostrarmos só como ponteiro então exibiremo o endereço da memória
	fmt.Println(*ponteiro)

	// podemos tambem atualizar o valor deste endereço na memória usando o ponteiro
	var ponteiroAtribuido *int = &a
	// uma vez atribuido usando * para dizer que esse endereço na memória agora passara a ter 50 como valor guardado
	*ponteiroAtribuido = 50

	// exibimos que a memória que armazena o valor não mudou (ja q armazenamos o mesmo endereço)
	fmt.Printf("Endereço de memória: %v\n", ponteiroAtribuido)
	// Mas mostramos que o valor armazenado nesse registro de memória foi atualziado
	fmt.Println(*ponteiroAtribuido)

	//Agora se recuperarmos o valor de &a ele ira mostrar tambem 50 pq o valor na memória foi mudado
	fmt.Println("Valor da variavel a",a)

}
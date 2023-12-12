package main

import (
	"fmt"
	"time"
)

func contador(tipo string){
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		//colocamos um tempo para que a rotina consiga ser processada do contrario como todos os dados são mais rapidos q a rotina go ele não sera exibido
		time.Sleep(time.Second)
	}
}

func main(){

	// agora usamos go para que ele crie uma thread e assim consigamos trabalhar com paralelismo e concorrencia.
	go contador("sem go routine")
	go contador("com go routine")

	//mesmo passando para a função nós ainda precisamos passar um sleep na função ou o valor não sera printado
	time.Sleep(time.Second * 10)
	

}
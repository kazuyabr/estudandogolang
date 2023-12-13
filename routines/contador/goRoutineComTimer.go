package main

import (
	"fmt"
	"time"
)

func contador(tipo string){
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}

func main(){
	contador("sem go routine")

	//agora iniciamos uma thread
	go contador("com go routine")
	fmt.Println("E ae 1")
	fmt.Println("E ae 2")

	//colocamos um tempo para que a rotina consiga ser processada do contrario como todos os dados são mais rapidos q a rotina go ele não sera exibido
	time.Sleep(time.Second)

}
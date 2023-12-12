package main

import (
	"fmt"
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
	
	//Exibimos os demais
	fmt.Println("E ae 1")
	fmt.Println("E ae 2")

	// se não colocamos um tempo para que a rotina consiga ser processada ela sera ignorada
	

}
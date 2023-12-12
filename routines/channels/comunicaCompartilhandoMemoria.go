package main

import "fmt"

//Entenda MAIN como uma thread
//T1
func main(){

	//Criamos o canal de comunicação entre as threads usando make()
	// T1 comunica com T2
	hello := make(chan string)

	//Entenda go como outra thread
	//T2
	go func(){
		hello <- "Hello World" //usa-se o simbolo de apontamento verso (<-) para dizer que o canal recebera o valor a seguir (este tambem é um simbulo de atribuição)
	}()

	result := <-hello
	fmt.Println(result)

}
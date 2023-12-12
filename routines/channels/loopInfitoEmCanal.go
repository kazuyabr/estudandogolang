package main

import "fmt"

func main(){

	forever := make(chan string)

	//neste caso esse processo ainda seria parado
	go func(){
		x := true
		for {
			if x == true{
				continue
			}
		}
	}()

	fmt.Println("Aguardando para sempre!")

	// mas como a thread 1 possui um valor que esta esperando alguem atribuir a ele, ele ira travar a execução para sempre
	<- forever

}
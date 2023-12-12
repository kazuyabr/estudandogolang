package main

import "fmt"

func main(){

	resultado, str := soma(10, 20)
	fmt.Println(resultado, str)

}

func soma(a int, b int) (int, string){
	return a + b, "somou!"
}
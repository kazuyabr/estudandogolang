package main

import (
	"fmt"
)

func main(){

	resultado := somaTudo(10,20,30,40,50)

	fmt.Println(resultado)
}

func somaTudo(x ...int) int {
	resultado := 0

	for _, v:=range x{
		resultado += v
	}

	return resultado
}
package main

import (
	"fmt"

	"curso-go/math"
)


func main(){

	resultado := math.Soma(1, 1)
	fmt.Printf("%v\n", resultado)
	fmt.Printf("%T\n", resultado)

	resultadoX := math.SomaX(1)
	fmt.Printf("%v\n", resultadoX)
	fmt.Printf("%T\n", resultadoX)

}

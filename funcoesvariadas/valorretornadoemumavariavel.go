package main

import (
	"fmt"
)

func main(){
	resultado := soma(10, 20)

	fmt.Println(resultado)
}

// Por padrão o restorno quando se tem a variavel, ele a identifica e retorna não precisando passar nada para return, só necessitando o próprio return
func soma(a int, b int) (result int){
	result = a + b
	return
}
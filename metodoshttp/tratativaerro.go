package main

import (
	"errors"
	"fmt"
	"log"
)

func main(){

	// Você pode bypassar erros usando _ ao invés de vc tratar
	// res, _:= soma(10, 10) isso se chama blanked undentified com isso a variavel de erro é jogada fora e se o erro ocorrer a verivel retornara apenas 0
	// no go você não pode deixar de receber valores

	res, err:= soma(10, 10)

	if err != nil{
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func soma(x int, y int) (int, error){
	res := x+y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}
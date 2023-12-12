package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// O que faremos neste struct é usar de tags para dizer ao json que ele deve exibir esses campos de outra maneira
type Pessoa struct {
	Nome string `json:"nome"`
	Cidade string `json:"cidade"`
	Pais string `json:"pais"`
}


func main(){

	fulano := Pessoa{
		Nome: "Fulano de tal",
		Cidade: "Onde judas perdeu as botas",
		Pais: "USA",
	}

	fmt.Printf("Sou %s e moro %s :D\n",fulano.Nome, fulano.Cidade)

	//Agora em json
	fulanoJSON, err := json.Marshal(fulano)
	if err != nil{
		log.Fatal(err.Error())
	}

	// bytes
	fmt.Println(fulanoJSON)

	// string
	fmt.Println(string(fulanoJSON))
}
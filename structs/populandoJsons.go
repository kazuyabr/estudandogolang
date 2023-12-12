package main

import (
	"encoding/json"
	"fmt"
)

// O que faremos neste struct é usar de tags para dizer ao json que ele deve exibir esses campos de outra maneira
type Pessoa struct {
	Nome string `json:"nome"`
	Cidade string `json:"cidade"`
	Pais string `json:"pais"`
}


func main(){
	//crases tambem são conhecidos como backtics
	// aqui irei criar manualmente um objeto json (na vdd uma string)
	jsonCliente := `{"nome":"Fulano de tal","cidade":"Onde judas perdeu as botas","pais":"USA"}`
	fulano := Pessoa{}

	fmt.Println("JSON String: ",jsonCliente)
	// Abaixo é mostrado fulano vazio mesmo (é para demonstrar que não foi instanciado)
	fmt.Printf("Sou %s e moro %s :D\n",fulano.Nome, fulano.Cidade)

	//Agora tratamos o json
	json.Unmarshal([]byte(jsonCliente), fulano)
	fmt.Println("Fulano não alterado: ", fulano)

	//Ponto importante, aqui nós apontamos os bytes de jsonCliente para uma referencia de memória para fulano, isso faz com que a referencia seja de fato auterada, pois do contrario só teria o valor neste escopo e o resultado seria inalterado.
	json.Unmarshal([]byte(jsonCliente), &fulano)
	fmt.Println("Fulano alterado por conta de termos apontado o endereço de memória: ", fulano)

	fmt.Printf("Sou %s e moro %s :D\n",fulano.Nome, fulano.Cidade)

}
package main

// Criamos um serviço basico onde faremos uma requisição e ela pode ser cancelada pelo cliente a qualquer momento
// se vc usa CURL ou cliente REST como Postman faça a requisição GET para localhost:8080 e depois a cancele antes de processar
// O resultado é que o servidor é inteligente para parar o processo se o cliente cancelar e assim ele não processa inutilmente ja q não tera para quem devolver o resultado

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Iniciamos o servidor na porta 8080
func main(){
	// Definimos uma rota simples sem muito tratamento
	http.HandleFunc("/", home)
	// Abrimos a porta 8080 para que espere requisições e passamos nil por padrão, mas pode ser tratado depois em caso de possiveis erros
	err := http.ListenAndServe(":8080", nil)

	// Tratamos possiveis erros
	if err != nil {
		fmt.Printf("could not serve on port 8080: %v\n", err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log.Println("Inicio minha request")
	defer log.Println("Encerrou minha request")

	select{
	case <-time.After(time.Second * 5):
		fmt.Fprintln(w, "Requisição processada com sucesso")
		// Quem fez a requisição ira receber como retorno a mensagem abaixo
		w.Write([]byte("Requisição processada com sucesso"))
	case <- ctx.Done():
		// caso o requisitate cancele manualmente a solicitação
		log.Println("Request Cancelada")
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)

	}
}
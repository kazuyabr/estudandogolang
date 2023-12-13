package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){
	//Criamos o contexto e ele é sempre o primeiro parametro de suas funções.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)	
	defer cancel()

	// Criamos uma requisição GET
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	// Tratamos possiveis erros
	if err != nil {
		log.Fatalf("Erro ao criar request: %v", err)
	}

	// Executamos a request
	res, err := http.DefaultClient.Do(req)

	// Tratamos possiveis erros
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Encerramos tudo
	defer res.Body.Close()

	// Exibimos o resultado
	io.Copy(os.Stdout, res.Body)
}
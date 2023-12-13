package main

import (
	"context"
	"fmt"
)

func main() {
    // Contexto de background (contexto em branco ou contexto raiz ou contexto pai, chame como achei melhor)
    ctx := context.Background()

    // Adicionamos uma chave-valor ao contexto
    ctx = context.WithValue(ctx, "nomeDoCliente", "João da Silva")

    // Imprimimos o valor da chave "nomeDoCliente"
    fmt.Println(ctx.Value("nomeDoCliente"))
}
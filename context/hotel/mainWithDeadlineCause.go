package main

// Como é uma adição recente a biblioteca golang 1.21 o contexto WithDeadlineCause ainda possui bugs ao emitir a causa.

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
  // Cria um contexto de background
  ctx := context.Background()

  // Cria um deadline e uma causa de cancelamento
  deadline := time.Now().Add(time.Second * 2)
  cause := errors.New("A operação demorou muito e foi cancelada pelo deadline")

  // Cria um novo contexto com deadline e causa
  newCtx, cancel := context.WithDeadlineCause(ctx, deadline, cause)

  // Usa o novo contexto para realizar sua operação
  go doSomething(newCtx)

  // Cancela o contexto caso necessário
  defer cancel()

  // Espera a operação terminar
  <-newCtx.Done()
}

func doSomething(ctx context.Context) {
  // Realiza a operação
  fmt.Println("A operação começou")
  time.Sleep(time.Second * 5)
  fmt.Println("A operação terminou")

  // Verifica se o contexto foi cancelado
  if ctx.Err() != nil {
    // O contexto foi cancelado
    fmt.Println("O contexto foi cancelado com a causa:", ctx.Err().Error())
  }
}

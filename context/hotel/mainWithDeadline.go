package main

import (
	"context"
	"fmt"
	"time"
)

// O exemplo desta aplicação é de um site de reservas de hotel como hotels.com ou booking.com por exemplo.

// Contextos podem ser comparados a arvores dentro do go onde vc tem contextos pai, contextos filhos e assim por diante
func main() {
    // Contexto de background (contexto em branco ou contexto raiz ou contexto pai, chame como achei melhor)
    ctx := context.Background()

    // Neste contexto nós configuramos uma possibilidade de o contexto ser cancelado
    // A Deadline faz com que depois do tempo limite ele encerre a operação
    ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*5))

    // Aqui garantimos que tudo será executado antes de rodar o cancel()
    // o defer faz isso, é como uma instrução que espera tudo passar para que no final ele seja executado
    defer cancel()

    // Aqui chamamos a função bookHotel, passando o contexto criado
    bookHotel(ctx)
}

// Por boas práticas é sempre recomendável que declare por primeiro o contexto a ser empregado na função
func bookHotel(ctx context.Context) {
    // Condição select considera todos os canais esperando que qualquer um de certo (pode ser comparado a switch case, mas seus cases são canais e não comparações)
    // Ou seja cada caso trabalhará paralelamente podendo sofrer com concorrência e cooperativismo

    // Aqui adicionamos uma nova linha de código
    select {
    case <-ctx.Done():
        // Se a operação for encerrada abruptamente então printa o erro
        fmt.Println(ctx.Err())
    case <-time.After(time.Second * 10):
        // Tempo para processar algo
        fmt.Println("Pronto!")
	}
}
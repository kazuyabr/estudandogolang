package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Criamos o contexto pai
    ctx := context.Background()

    // Preparamos um timeout para esse contexto
    ctx, cancel := context.WithTimeout(ctx, time.Second*10)

    // Aqui garantimos que tudo será executado antes de rodar o cancel()
    // o defer faz isso, é como uma instrução que espera tudo passar para que no final ele seja executado
    defer cancel()

    // Aqui chamamos a função bookHotel, passando o contexto criado
    bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
    // Aqui começamos a operação de reservar um hotel
    fmt.Println("Iniciando a operação de reserva de hotel...")

    // Aqui fazemos uma chamada a uma API externa para reservar o hotel
    // Essa chamada pode demorar alguns segundos
    time.Sleep(time.Second * 5)

    // Aqui verificamos se o contexto foi cancelado
    if ctx.Err() != nil {
        // Se o contexto foi cancelado, imprimimos uma mensagem de erro
        fmt.Println("A operação de reserva de hotel foi cancelada!")
        return
    }

    // Aqui imprimimos uma mensagem de sucesso
    fmt.Println("Reserva de hotel realizada com sucesso!")

    // Aqui aplicamos o conhecimento sobre context.TODO()
    // Usamos context.TODO() como o contexto para a chamada a API externa
	// Este contexto impede que o mesmo seja encerrado uma vez que o contexto não foi especificado impedindo que o contexto seja encerrado.
    ctx = context.TODO()

    // Aqui fazemos uma nova chamada a API externa para reservar o hotel
    // Essa chamada pode demorar alguns segundos
    time.Sleep(time.Second * 5)

    // Aqui verificamos se o contexto foi cancelado
    if ctx.Err() != nil {
        // Se o contexto foi cancelado, imprimimos uma mensagem de erro
        fmt.Println("A operação de reserva de hotel foi cancelada!")
        return
    }

    // Aqui imprimimos uma mensagem de sucesso
    f
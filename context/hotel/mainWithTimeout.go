package main

import (
	"context"
	"time"
)

// O exemplo desta aplicação é de um site de reservas de hotel como hotels.com ou booking.com por exemplo.

// Contextos podem ser comparados a arvores dentro do go onde vc tem contextos pai, contextos filhos e assim por diante
func main(){
	// Contexto de background (contexto em branco ou contexto raiz ou contexto pai, chame como achei melhor)
	ctx := context.Background()
	// Neste contexto nós configuramos uma possibilidade de o contexto ser cancelado
	// Por boas praticas, vc sempre cancelara o contexto no final a menos que o contexto seja cancelado antes
	ctx, cancel := context.WithTimeout(ctx, time.Second*10) //WithTimeout garante que o contexto seja encerrado após um periodo de tempo

	// Aqui garantimos que tudo sera executado antes de rodar o cancel() o defer faz isso é como uma instrução que espera tudo passar para que no final ele seja executado
	defer cancel()


	bookHotel(ctx)

}

// Por boas praticas é sempre recomendavel que declare por primeiro o contexto a ser empregado na função
func bookHotel(ctx context.Context){
	// Condição select considera todos os canais esperando que qualquer um de certo (pode ser comparado a switch case, mas seus cases são canais e nao comparações)
	// Ou seja cada caso trabalhara paralelamente podendo sofrer com concorrencia e cooperativismo

	select{
	case <- ctx.Done():
		// Se o processo for cancelado antes de 5 segundos o quarto não podera ser reservado
		println("Tempo excedido para reservar o quarto")
	case <-time.After(time.Second * 5):
		// Se o processo passo o tempo configurado ele ira reservar com sucesso esse quarto
		println("Quarto reservado com sucesso")
	}
}
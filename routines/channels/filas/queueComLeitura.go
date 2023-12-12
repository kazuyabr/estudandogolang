package main

import (
	"fmt"
	"time"
)

func main(){
	queue := make(chan int)

	go func(){
		i := 0

		for {
			time.Sleep(time.Second)
			queue <- i
			i++
		}
	}()

	//o loop só vai parar quando a fila for esvaziada. então se queue não tiver mais valores o processo é encerrado
	for x := range queue{
		fmt.Println(x)
	}
}
package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int){
	for res := range msg{
		fmt.Println("Worker: ", workerId," Msg: ", res)
		time.Sleep(time.Second)
	}
}

func main(){

	msg := make(chan int)

	// cada route só usa 2k ao inves de 1m das threads do SO ja que faz parte da própria runtime. Go não chama threads do SO e sim de sua runtime o que poupa muito recurso
	go worker(1, msg) 
	go worker(2, msg)
	go worker(3, msg)
	go worker(4, msg)

	for i := 0; i <10; i++{
		msg <- i
	}

}
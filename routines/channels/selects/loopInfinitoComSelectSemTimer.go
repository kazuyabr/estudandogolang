package main

import (
	"fmt"
	"time"
)

func main(){
	hello := make(chan string)

	go func(){
		time.Sleep(time.Second * 2)
		hello <- "hello world"
	}()

	select {
	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("default")
	}

}
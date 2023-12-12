package main

import "time"

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
}
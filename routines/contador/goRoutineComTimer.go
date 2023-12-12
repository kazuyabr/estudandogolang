package main

import (
	"fmt"
	"time"
)

func contador(tipo string){
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}

func main(){
	contador("sem go routine")
	go contador("com go routine")
	fmt.Println("E ae 1")
	fmt.Println("E ae 2")
	time.Sleep(time.Second)

}
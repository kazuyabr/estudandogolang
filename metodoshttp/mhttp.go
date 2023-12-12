package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	res, err := http.Get("http://google.com.br")
	
	if err != nil{
		log.Fatal(err.Error())
	}
	
	fmt.Printf("%+v\n", res.Header)
}
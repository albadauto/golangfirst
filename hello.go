package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com.br") //Requisição get no site da google
	if err != nil {
		log.Fatal("Deu ruim amigão", err.Error())
	}

	fmt.Println(res.Header)

}


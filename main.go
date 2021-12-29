package main

import (
	"log"

	"github.com/darias-developer/challenge-go/config"
	"github.com/darias-developer/challenge-go/handler"
)

func main() {

	if config.CheckConnection() == 0 {
		log.Fatal("Error on db conexion")
		return
	}

	handler.RouterManager()
}

package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/darias-developer/challenge-go/config"
	"github.com/darias-developer/challenge-go/handler"
	"github.com/darias-developer/challenge-go/middleware"
)

func main() {

	//carga variables desde archivo .env
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//inicializa la configuracion de los logs
	middleware.LoggerInit()

	middleware.LogInfo.Println("init main")

	//crea un ping a la db
	if config.CheckConnection() == 0 {
		log.Fatal("Error on db conexion")
		return
	}

	middleware.LogInfo.Println("Conexion success to DB!!")

	//carga las rutas
	handler.RouterManager()
}

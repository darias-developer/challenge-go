package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/darias-developer/challenge-go/middleware"
	"github.com/darias-developer/challenge-go/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* RouterManager maneja las rutas y puertos */
func RouterManager() {

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/signup", middleware.CheckDB(router.Signup)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(newRouter)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

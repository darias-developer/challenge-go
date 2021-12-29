package middleware

import (
	"net/http"

	"github.com/darias-developer/challenge-go/config"
)

/* CheckDB valida la conexion a la db antes de llamar a una funcion */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {
		if config.CheckConnection() == 0 {
			http.Error(rw, "Conexion lost in db", 500)
		}

		next.ServeHTTP(rw, r)
	}
}

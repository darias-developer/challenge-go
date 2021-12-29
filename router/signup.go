package router

import (
	"encoding/json"
	"net/http"

	"github.com/darias-developer/challenge-go/model"
	"github.com/darias-developer/challenge-go/service"
)

/* Signup router para la creacion de usuarios */
func Signup(rw http.ResponseWriter, r *http.Request) {

	var userModel model.UserModel
	err := json.NewDecoder(r.Body).Decode(&userModel)

	if err != nil {
		http.Error(rw, "Error in the data: "+err.Error(), 400)
		return
	}

	if len(userModel.Email) == 0 {
		http.Error(rw, "El email del usuario es requerido", 400)
		return
	}

	if len(userModel.Password) < 6 {
		http.Error(rw, "La password tiene un largo no valido. (min 6)", 400)
		return
	}

	_, isFound, _ := service.FindUserByEmail(userModel.Email)

	if isFound {
		http.Error(rw, "el Usuario ya existe", 400)
		return
	}

	_, status, err := service.CreateUser(userModel)

	if err != nil {
		http.Error(rw, "Ha ocurrido un error al crear el usuario: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(rw, "Ha ocurrido un error al crear el usuario en la db: "+err.Error(), 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

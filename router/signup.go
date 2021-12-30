package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/darias-developer/challenge-go/data"
	"github.com/darias-developer/challenge-go/middleware"
	"github.com/darias-developer/challenge-go/model"
	"github.com/darias-developer/challenge-go/service"
)

/* Signup router para la creacion de usuarios */
func Signup(rw http.ResponseWriter, r *http.Request) {

	middleware.LogInfo.Println("init Signup")

	err := process(r)

	var status int
	var response data.Response

	if err != nil {

		middleware.LogError.Println(err.Error())

		response = data.Response{
			ResponseCode: "ERROR",
			Description:  err.Error(),
		}

		status = http.StatusBadRequest
	} else {

		response = data.Response{
			ResponseCode: "SUCCESS",
			Description:  "El usuario se ha creado de forma existosa",
		}

		status = http.StatusCreated
	}

	responseJson, _ := json.Marshal(response)

	middleware.LogInfo.Println(string(responseJson))
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(responseJson)

	middleware.LogInfo.Println("end Signup")
}

func process(r *http.Request) error {

	var userModel model.UserModel
	err := json.NewDecoder(r.Body).Decode(&userModel)

	if err != nil {
		middleware.LogError.Println(err)
		return errors.New("Existe un error en la data enviada")
	}

	if len(userModel.Email) == 0 {
		return errors.New("El email del usuario es requerido")
	}

	if len(userModel.Password) < 6 {
		return errors.New("La password tiene un largo no valido. (min 6)")
	}

	_, isFound, _ := service.FindUserByEmail(userModel.Email)

	if isFound {
		return errors.New("El Usuario ya existe")
	}

	_, status, err := service.CreateUser(userModel)

	if err != nil {
		middleware.LogError.Println(err)
		return errors.New("Ha ocurrido un error al crear el usuario")
	}

	if !status {
		middleware.LogError.Println(err)
		return errors.New("Ha ocurrido un error al crear el usuario en la db")
	}

	return nil
}

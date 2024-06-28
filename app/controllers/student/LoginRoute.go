package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	requests "github.com/fateqense/siga/app/requests/student"
	usecases "github.com/fateqense/siga/app/usecases/student"
	"github.com/fateqense/siga/utils"
)

func (StudentController) LoginRoute(w http.ResponseWriter, r *http.Request) {
	var loginRequest requests.LoginRequest
	if err := utils.DecodeJSONBody(w, r, &loginRequest); err != nil {
		var malformedRequest *utils.MalformedRequest

		if errors.As(err, &malformedRequest) {
			http.Error(w, malformedRequest.Error(), malformedRequest.Status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	studentUseCase := usecases.StudentUseCase{}

	res, err := studentUseCase.LoginAction(loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(token)
}

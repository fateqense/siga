package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	usecases "github.com/fateqense/siga/app/usecases/student"
)

func (StudentController) GetHistoryRoute(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("authorization")
	if authorization == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session := strings.Split(authorization, " ")[1]

	studentUseCase := usecases.StudentUseCase{}

	completeHistory, err := studentUseCase.GetHistoryAction(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	completeHistoryJson, err := json.Marshal(completeHistory)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(completeHistoryJson)
}

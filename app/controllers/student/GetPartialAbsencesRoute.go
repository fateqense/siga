package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	usecases "github.com/fateqense/siga/app/usecases/student"
)

func (StudentController) GetPartialAbsencesRoute(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("authorization")
	if authorization == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session := strings.Split(authorization, " ")[1]

	studentUseCase := usecases.StudentUseCase{}

	partialAbsences, err := studentUseCase.GetPartialAbsencesAction(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	partialAbsencesJson, err := json.Marshal(partialAbsences)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(partialAbsencesJson)
}

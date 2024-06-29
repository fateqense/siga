package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	usecases "github.com/fateqense/siga/app/usecases/student"
)

func (StudentController) GetPartialGradesRoute(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("authorization")
	if authorization == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session := strings.Split(authorization, " ")[1]

	studentUseCase := usecases.StudentUseCase{}

	partialGrades, err := studentUseCase.GetPartialGradesAction(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	partialGradesJson, err := json.Marshal(partialGrades)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(partialGradesJson)
}

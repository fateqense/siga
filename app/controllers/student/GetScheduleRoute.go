package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	usecases "github.com/fateqense/siga/app/usecases/student"
)

func (StudentController) GetScheduleRoute(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("authorization")
	if authorization == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session := strings.Split(authorization, " ")[1]

	studentUseCase := usecases.StudentUseCase{}

	schedule, err := studentUseCase.GetScheduleAction(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scheduleJson, err := json.Marshal(schedule)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(scheduleJson)
}

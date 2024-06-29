package server

import (
	"net/http"

	controllers "github.com/fateqense/siga/app/controllers/student"
)

func (Server) BuildRoutes() *http.ServeMux {
	api := http.NewServeMux()

	studentController := controllers.StudentController{}
	api.HandleFunc("POST /auth/login", http.HandlerFunc(studentController.LoginRoute))
	api.HandleFunc("GET /profile", http.HandlerFunc(studentController.GetProfileRoute))
	api.HandleFunc("GET /grades", http.HandlerFunc(studentController.GetPartialGradesRoute))

	return api
}

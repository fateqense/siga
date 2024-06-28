package server

import "net/http"

func (Server) BuildRoutes() *http.ServeMux {
	api := http.NewServeMux()

	return api
}

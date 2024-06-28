package server

import (
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	addr string
}

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	srv := &Server{
		addr: fmt.Sprintf(":%s", port),
	}

	return &http.Server{
		Addr:    srv.addr,
		Handler: srv.BuildRoutes(),
	}
}

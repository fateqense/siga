package main

import (
	"fmt"

	"github.com/fateqense/siga/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	srv := server.NewServer()

	fmt.Printf("🎒 Server is running at %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

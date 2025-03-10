package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/ioansx/clientele/internal/routes"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	serveMux := http.NewServeMux()
	routes.AddRoutes(serveMux)

	addr := net.JoinHostPort("", port)
	slog.Info(fmt.Sprintf("Listening on %s...", addr))
	if err := http.ListenAndServe(addr, serveMux); err != nil {
		log.Fatal(err)
	}
}

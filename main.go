package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bjcorder/go-api-boilerplate/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Initializing server...")

	godotenv.Load(".env")

	portString := os.Getenv("SERVER_PORT")
	if portString == "" {
		log.Fatal("Unable to get value: SERVER_PORT")
	}

	router := chi.NewRouter()

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerHealthz)
	v1Router.Get("/error", handlerError)

	router.Mount("/v1", v1Router)

	serve := &http.Server{
		Handler: stack(router),
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

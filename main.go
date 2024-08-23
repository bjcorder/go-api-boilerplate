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

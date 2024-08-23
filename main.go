package main

import (
	"fmt"
	"net/http"

	"github.com/bjcorder/go-api-boilerplate/middleware"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/test", basicHandler)

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":3000",
		Handler: stack(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Test"))

}

package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Conent-Type", "text/plain")
		w.Write([]byte("Hello i am apigateway"))
	})

	




	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error open api_gateway http localhost:8080")
	}

}

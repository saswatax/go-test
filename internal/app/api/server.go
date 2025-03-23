package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer() {
	port := "8080"

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running"))
	})

	router := createRouter()
	r.Mount("/v1", router)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Printf("Server failed to start with err: %v", err)
		return
	}

	fmt.Printf("Server running on port: %v", port)
}

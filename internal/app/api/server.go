package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func StartServer() {
	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	r.Use(testMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		message := r.Context().Value("message").(string)

		fmt.Println(message)

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

package api

import "github.com/go-chi/chi/v5"

func createRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{id}", getPost)
	r.Get("/", getAllPost)
	r.Post("/", createPost)
	r.Patch("/", updatePost)
	r.Put("/", replacePost)
	r.Delete("/", deletePost)

	return r
}

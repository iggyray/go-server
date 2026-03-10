package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	controller "github.com/iggyray/go-server/internal/http/controller"
	handler "github.com/iggyray/go-server/internal/http/handler"
	"github.com/iggyray/go-server/internal/service"
)

const (
	PORT = ":8000"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is online!! 🚀🚀🚀"))
	})

	postService := service.New()
	postController := controller.New(postService)
	postHandler := handler.New(postController)
	r.Get("/posts", postHandler.GetPosts)
	r.Get("/posts/{id}", postHandler.GetPost)

	log.Printf("Listening on http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}

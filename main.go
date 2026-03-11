package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iggyray/go-server/internal/http/controller"
	"github.com/iggyray/go-server/internal/http/handler"
	"github.com/iggyray/go-server/internal/service"
	"github.com/lpernett/godotenv"
)

const (
	PORT = ":8000"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

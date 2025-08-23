package router

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iggyray/go-server/internal/service"
)

func PostRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getPosts)
	r.Get("/{id}", getPost)

	return r
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	postService := service.NewPostService()
	post := postService.GetPost()
	log.Printf("%+v", post)

	_, err := io.WriteString(w, fmt.Sprintf("%+v", post))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	postService := service.NewPostService()
	post := postService.GetPostById(postId)
	log.Printf("%+v", post)

	_, err := io.WriteString(w, fmt.Sprintf("%+v", post))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

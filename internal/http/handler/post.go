package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := h.controller.GetPosts()
	log.Printf("%+v", posts)

	_, err := io.WriteString(w, fmt.Sprintf("%+v", posts))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	post := h.controller.GetPost(postId)
	log.Printf("%+v", post)

	_, err := io.WriteString(w, fmt.Sprintf("%+v", post))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

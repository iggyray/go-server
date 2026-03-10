package handler

import (
	"github.com/iggyray/go-server/internal/domain"
)

type (
	PostHandler struct {
		controller PostController
	}

	PostController interface {
		GetPosts() []domain.Post
		GetPost(string) domain.Post
	}
)

func New(controller PostController) *PostHandler {
	return &PostHandler{controller: controller}
}

package controller

import (
	"github.com/iggyray/go-server/internal/domain"
)

type (
	PostController struct {
		service PostService
	}

	PostService interface {
		GetPosts() []domain.Post
		GetPost(string) domain.Post
	}
)

func New(service PostService) *PostController {
	return &PostController{service: service}
}

func (c *PostController) GetPosts() []domain.Post {
	return c.service.GetPosts()
}

func (c *PostController) GetPost(id string) domain.Post {
	return c.service.GetPost(id)
}

package service

import (
	"fmt"

	"github.com/iggyray/go-server/internal/domain"
)

type PostService struct{}

func NewPostService() PostService {
	return PostService{}
}

func (ps PostService) GetPost() domain.Post {
	post := domain.GenPost()

	return post
}

func (ps PostService) GetPostById(id string) domain.Post {
	post := domain.GenPost()
	post.Content = fmt.Sprintf("Post with post id: %s!", id)

	return post
}

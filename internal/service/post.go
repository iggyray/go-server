package service

import (
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

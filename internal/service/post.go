package service

import (
	"fmt"
	"strconv"

	"github.com/iggyray/go-server/internal/domain"
)

type PostService struct{}

func New() *PostService {
	return &PostService{}
}

func genPost(id string) domain.Post {
	return domain.Post{ID: domain.PostID(id), Author: "author-" + id, Content: id}
}

func genPosts(length int) []domain.Post {
	out := make([]domain.Post, length)
	for i := range out {
		out[i] = genPost(strconv.Itoa(i))
	}
	return out
}

func (p *PostService) GetPosts() []domain.Post {
	return genPosts(5)
}

func (p *PostService) GetPost(id string) domain.Post {
	post := genPost(id)
	post.Content = fmt.Sprintf("Post with post id: %s!", id)

	return post
}

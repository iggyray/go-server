package controller

import (
	"fmt"
	"strconv"

	"github.com/iggyray/go-server/internal/domain"
)

type PostController struct{}

func New() *PostController {
	return &PostController{}
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

func (p PostController) GetPosts() []domain.Post {
	return genPosts(5)
}

func (p PostController) GetPost(id string) domain.Post {
	post := genPost(id)
	post.Content = fmt.Sprintf("Post with post id: %s!", id)

	return post
}

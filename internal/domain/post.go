package domain

type (
	PostID string

	Post struct {
		ID      PostID
		Author  string
		Content string
	}
)

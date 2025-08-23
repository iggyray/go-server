package domain

type Post struct {
	Author  string
	Content string
}

func GenPost() Post {
	return Post{Author: "Bob", Content: "Hello world"}
}

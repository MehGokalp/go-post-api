package comment

import "github.com/go-post-api/pkg/post"

type Comment struct {
	Id         int
	Post       post.Post
	AuthorName string
	Email      string
	Body       string
}

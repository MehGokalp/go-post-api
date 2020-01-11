package post

import (
	"github.com/go-post-api/pkg/user"
	"time"
)

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	User      *user.User
	CreatedAt time.Time
}

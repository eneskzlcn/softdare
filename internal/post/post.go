package post

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"strings"
	"time"
)

type CreatePostRequest struct {
	ID      string
	UserID  string
	Content string
}
type CreatePostInput struct {
	Content string
}

func (c *CreatePostInput) Prepare() {
	c.Content = strings.TrimSpace(c.Content)
	c.Content = strings.ReplaceAll(c.Content, "\n\n", "\n")
	c.Content = strings.ReplaceAll(c.Content, "  ", " ")
}

func (c *CreatePostInput) Validate() error {
	return validation.Validate(c.Content, validation.Length(2, 1000))
}

type CreatePostResponse struct {
	ID        string
	CreatedAt time.Time
}

const userContextKey = "user"

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Post struct {
	ID        string
	UserID    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
}

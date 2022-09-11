package post

import (
	postUtil "github.com/eneskzlcn/softdare/internal/util/post"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"html/template"
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
type FormattedPost struct {
	ID        string
	UserID    string
	Content   template.HTML
	CreatedAt string
	UpdatedAt string
	Username  string
}

func FormatPost(post *Post) (formattedPost FormattedPost) {
	formattedPost.Content = postUtil.FormatPostContent(post.Content)
	formattedPost.CreatedAt = postUtil.FormatPostTime(post.CreatedAt)
	formattedPost.UpdatedAt = postUtil.FormatPostTime(post.UpdatedAt)
	formattedPost.ID = post.ID
	formattedPost.UserID = post.UserID
	formattedPost.Username = post.Username
	return
}

type postData struct {
	Post FormattedPost
}

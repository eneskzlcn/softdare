package post

import (
	"github.com/eneskzlcn/softdare/internal/comment"
	"github.com/eneskzlcn/softdare/internal/util/convertion"
	postUtil "github.com/eneskzlcn/softdare/internal/util/post"
	sessionUtil "github.com/eneskzlcn/softdare/internal/util/session"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"net/url"
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
	ID           string
	UserID       string
	Content      string
	CommentCount int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Username     string
}
type FormattedPost struct {
	ID           string
	UserID       string
	Content      string
	CommentCount int
	CreatedAt    string
	UpdatedAt    string
	Username     string
}

func FormatPost(post *Post) (formattedPost FormattedPost) {
	formattedPost.CreatedAt = postUtil.FormatPostTime(post.CreatedAt)
	formattedPost.UpdatedAt = postUtil.FormatPostTime(post.UpdatedAt)
	formattedPost.ID = post.ID
	formattedPost.Content = post.Content
	formattedPost.CommentCount = post.CommentCount
	formattedPost.UserID = post.UserID
	formattedPost.Username = post.Username
	return
}

type postData struct {
	Session  SessionData        `json:"session"`
	Post     FormattedPost      `json:"post"`
	Comments []FormattedComment `json:"comments"`
}
type SessionData struct {
	IsLoggedIn         bool
	User               sessionUtil.UserSessionData
	CreateCommentForm  url.Values
	CreateCommentError error
}
type FormattedComment struct {
	ID        string
	PostID    string
	Content   string
	UserID    string
	Username  string
	CreatedAt string
	UpdatedAt string
}

func FormatComments(comments []*comment.Comment) []FormattedComment {
	formattedComments := make([]FormattedComment, 0)
	for _, comment := range comments {
		formattedComments = append(formattedComments, formatComment(comment))
	}
	return formattedComments
}
func formatComment(comment *comment.Comment) (formattedComment FormattedComment) {
	formattedComment.ID = comment.ID
	formattedComment.PostID = comment.PostID
	formattedComment.Username = comment.Username
	formattedComment.Content = comment.Content
	formattedComment.CreatedAt = postUtil.FormatPostTime(comment.CreatedAt)
	formattedComment.UpdatedAt = postUtil.FormatPostTime(comment.UpdatedAt)
	return
}
func sessionDataFromRequest(session SessionProvider, r *http.Request) SessionData {
	var out SessionData
	generalSession := sessionUtil.GeneralSessionDataFromRequest(session, r)
	if generalSession.IsLoggedIn {
		out.IsLoggedIn = generalSession.IsLoggedIn
		out.User.ID = generalSession.User.ID
		out.User.Email = generalSession.User.Email
		out.User.Username = generalSession.User.Username
	}
	if session.Exists(r, "create-comment-error") {
		out.CreateCommentError = session.PopError(r, "create-comment-error")
	}
	if session.Exists(r, "create-comment-form") {
		form := session.Get(r, "create-comment-form")
		urlForm, err := convertion.AnyToGivenType[url.Values](form)
		if err == nil {
			out.CreateCommentForm = urlForm
		}
	}
	return out
}

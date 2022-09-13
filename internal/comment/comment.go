package comment

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rs/xid"
	"strings"
)

type IncreasePostCommentCountMessage struct {
	PostID         string `json:"post_id"`
	IncreaseAmount int    `json:"increase_amount"`
}
type CreateCommentRequest struct {
	ID      string
	PostID  string
	UserID  string
	Content string
}
type CreateCommentInput struct {
	PostID  string
	Content string
}

func (c *CreateCommentInput) Prepare() {
	c.Content = strings.TrimSpace(c.Content)
	c.Content = strings.ReplaceAll(c.Content, "\n\n", "\n")
	c.Content = strings.ReplaceAll(c.Content, "  ", " ")
}
func (c *CreateCommentInput) Validate() error {
	_, err := xid.FromString(c.PostID)
	if err != nil {
		return entity.InvalidPostID.Err()
	}
	err = validation.Validate(c.Content, validation.Length(2, 1000))
	if err != nil {
		return entity.InvalidContent.Err()
	}
	return nil
}

package commentutil_test

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/commentutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatComment(t *testing.T) {
	comment := &entity.Comment{
		ID:        "",
		PostID:    "",
		UserID:    "",
		Content:   "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "",
	}
	timeAgoFormatter := func(t time.Time) string {
		return "Just now"
	}
	formattedPost := commentutil.FormatComment(comment, timeAgoFormatter)
	assert.Equal(t, formattedPost.CreatedAt, "Just now")
	assert.Equal(t, formattedPost.UpdatedAt, "Just now")

}

func TestFormatComments(t *testing.T) {
	comments := []*entity.Comment{
		{
			ID:        "",
			PostID:    "",
			UserID:    "",
			Content:   "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Username:  "",
		},
		{
			ID:        "",
			PostID:    "",
			UserID:    "",
			Content:   "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Username:  "",
		},
		{
			ID:        "",
			PostID:    "",
			UserID:    "",
			Content:   "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Username:  "",
		},
	}
	timeAgoFormatter := func(t time.Time) string {
		return "Just now"
	}
	formattedComments := commentutil.FormatComments(comments, timeAgoFormatter)
	for _, comment := range formattedComments {
		assert.Equal(t, comment.CreatedAt, "Just now")
		assert.Equal(t, comment.UpdatedAt, "Just now")
	}
}

package postutil_test

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/postutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatPost(t *testing.T) {
	post := &entity.Post{
		ID:           "",
		UserID:       "",
		Content:      "",
		CommentCount: 0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Username:     "",
	}
	timeAgoFormatter := func(time.Time) string {
		return "just now"
	}
	formattedPost := postutil.FormatPost(post, timeAgoFormatter)
	assert.Equal(t, formattedPost.CreatedAt, "just now")
	assert.Equal(t, formattedPost.UpdatedAt, "just now")
}

func TestFormatPosts(t *testing.T) {
	posts := []*entity.Post{
		{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	timeAgoFormatter := func(time.Time) string {
		return "just now"
	}
	formattedPosts := postutil.FormatPosts(posts, timeAgoFormatter)
	for _, formattedPost := range formattedPosts {
		assert.Equal(t, formattedPost.CreatedAt, "just now")
		assert.Equal(t, formattedPost.UpdatedAt, "just now")
	}
}

package userutil_test

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/userutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatUserWithFollowingOption(t *testing.T) {
	givenUser := entity.UserWithFollowedOption{
		ID:            "testid",
		Email:         "testEmail",
		Username:      "testUsername",
		PostCount:     3,
		FollowerCount: 2,
		FollowedCount: 1,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		IsFollowed:    true,
	}
	formattedUser := userutil.FormatUserWithFollowingOption(&givenUser, func(t time.Time) string {
		return "a time ago"
	})
	assert.Equal(t, "a time ago", formattedUser.CreatedAt)
	assert.Equal(t, "a time ago", formattedUser.UpdatedAt)
	assert.Equal(t, givenUser.Username, formattedUser.Username)
	assert.Equal(t, givenUser.Email, formattedUser.Email)
	assert.Equal(t, givenUser.PostCount, formattedUser.PostCount)
	assert.Equal(t, givenUser.FollowerCount, formattedUser.FollowerCount)
	assert.Equal(t, givenUser.FollowedCount, formattedUser.FollowedCount)
	assert.Equal(t, givenUser.IsFollowed, formattedUser.IsFollowed)
}
func TestFormatUsersWithFollowingOption(t *testing.T) {
	givenUsers := []*entity.UserWithFollowedOption{
		{
			ID:            "testid",
			Email:         "testEmail",
			Username:      "testUsername",
			PostCount:     3,
			FollowerCount: 2,
			FollowedCount: 1,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			IsFollowed:    true,
		},
		{
			ID:            "testid",
			Email:         "testEmail",
			Username:      "testUsername",
			PostCount:     3,
			FollowerCount: 2,
			FollowedCount: 1,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			IsFollowed:    true,
		},
	}
	formattedUsers := userutil.FormatUsersWithFollowingOption(givenUsers, func(t time.Time) string {
		return "a time ago"
	})
	for i := 0; i < len(formattedUsers); i++ {
		assert.Equal(t, "a time ago", formattedUsers[i].CreatedAt)
		assert.Equal(t, "a time ago", formattedUsers[i].UpdatedAt)
		assert.Equal(t, givenUsers[i].Username, formattedUsers[i].Username)
		assert.Equal(t, givenUsers[i].Email, formattedUsers[i].Email)
		assert.Equal(t, givenUsers[i].PostCount, formattedUsers[i].PostCount)
		assert.Equal(t, givenUsers[i].FollowerCount, formattedUsers[i].FollowerCount)
		assert.Equal(t, givenUsers[i].FollowedCount, formattedUsers[i].FollowedCount)
		assert.Equal(t, givenUsers[i].IsFollowed, formattedUsers[i].IsFollowed)
	}

}

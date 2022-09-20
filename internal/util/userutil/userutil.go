package userutil

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	"time"
)

func FormatUserWithFollowingOption(user *entity.UserWithFollowedOption, timeFormatter func(time.Time) string) (formattedUser entity.FormattedUserWithFollowedOption) {
	formattedUser.ID = user.ID
	formattedUser.Email = user.Email
	formattedUser.Username = user.Username
	formattedUser.IsFollowed = user.IsFollowed
	formattedUser.FollowedCount = user.FollowedCount
	formattedUser.FollowerCount = user.FollowerCount
	formattedUser.PostCount = user.PostCount
	formattedUser.CreatedAt = timeFormatter(user.CreatedAt)
	formattedUser.UpdatedAt = timeFormatter(user.UpdatedAt)
	return
}

func FormatUsersWithFollowingOption(users []*entity.UserWithFollowedOption, timeFormatter func(time.Time) string) (formattedUsers []entity.FormattedUserWithFollowedOption) {
	for _, user := range users {
		formattedUsers = append(formattedUsers, FormatUserWithFollowingOption(user, timeFormatter))
	}
	return
}

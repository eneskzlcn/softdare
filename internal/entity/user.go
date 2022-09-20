package entity

import "time"

type User struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	PostCount     int       `json:"post_count"`
	FollowerCount int       `json:"follower_count"`
	FollowedCount int       `json:"followed_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserWithFollowedOption struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	PostCount     int       `json:"post_count"`
	FollowerCount int       `json:"follower_count"`
	FollowedCount int       `json:"followed_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	IsFollowed    bool      `json:"is_followed"`
}

type FormattedUserWithFollowedOption struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	PostCount     int    `json:"post_count"`
	FollowerCount int    `json:"follower_count"`
	FollowedCount int    `json:"followed_count"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	IsFollowed    bool   `json:"is_followed"`
}

/*UserIdentity keeps the values that provide user identify.
The fields can be change for usage like just with id or password.
Do not touch the name of the struct due to dependent parts. If you
change the struct name be sure to reformat related parts.

*/
type UserIdentity struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

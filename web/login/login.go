package login

import "time"

type LoginInput struct {
	Email    string
	Username *string
}

type CreateUserRequest struct {
	ID       string
	Email    string
	Username string
}

type User struct {
	ID        string
	Email     string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

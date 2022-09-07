package login

import (
	"context"
	"errors"
	"github.com/rs/xid"
	"time"
)

type LoginRepository interface {
	IsUserExistsByEmail(ctx context.Context, email string) (bool, error)
	IsUserExistsByUsername(ctx context.Context, username string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, request CreateUserRequest) (time.Time, error)
}
type Service struct {
	repository LoginRepository
}

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUsernameAlreadyTaken = errors.New("username already taken")
)

func NewService(repository LoginRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Login(ctx context.Context, inp LoginInput) (*User, error) {
	exists, err := s.repository.IsUserExistsByEmail(ctx, inp.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return s.repository.GetUserByEmail(ctx, inp.Email)
	}
	if inp.Username == nil {
		return nil, ErrUserNotFound
	}
	exists, err = s.repository.IsUserExistsByUsername(ctx, *inp.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUsernameAlreadyTaken
	}
	// if both username and email is given bot not found in db then create the user
	id := xid.New().String()
	createdAt, err := s.repository.CreateUser(ctx, CreateUserRequest{
		ID:       id,
		Email:    inp.Email,
		Username: *inp.Username,
	})
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        id,
		Email:     inp.Email,
		Username:  *inp.Username,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}, nil
}

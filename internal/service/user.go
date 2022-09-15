package service

import (
	"context"
	"errors"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/rs/xid"
	"time"
)

var ErrUsernameAlreadyTaken = errors.New("username already taken")

func (s *Service) Login(ctx context.Context, email string, username *string) (*entity.User, error) {
	if err := validation.IsValidEmail(email); err != nil {
		return nil, entity.InvalidEmail
	}
	exists, err := s.repository.IsUserExistsByEmail(ctx, email)
	if err != nil {
		s.logger.Error("error is user exists by email from repository")
		return nil, entity.UserNotFound
	}
	if exists {
		s.logger.Debug("user getting from repository")
		return s.repository.GetUserByEmail(ctx, email)
	}
	if username == nil || *username == "" {
		s.logger.Error("username not given")
		return nil, entity.UsernameNotGiven
	}
	exists, err = s.repository.IsUserExistsByUsername(ctx, *username)
	if err != nil {
		s.logger.Error("getting user exist by username from repository error")
		return nil, entity.UserNotFound
	}
	if exists {
		return nil, ErrUsernameAlreadyTaken
	}
	if err := validation.IsValidUsername(*username); err != nil {
		s.logger.Error("invalid username input", s.logger.StringModifier("username", *username))
		return nil, entity.InvalidUsername
	}
	id := xid.New().String()
	//if validated then create the user
	createdAt, err := s.repository.CreateUser(ctx, id, email, *username)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return &entity.User{
		ID:            id,
		Email:         email,
		PostCount:     0,
		FollowerCount: 0,
		FollowedCount: 0,
		Username:      *username,
		CreatedAt:     createdAt,
		UpdatedAt:     createdAt,
	}, nil
}
func (s *Service) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return s.repository.GetUserByUsername(ctx, username)
}
func (s *Service) IncreaseUserPostCount(ctx context.Context, userID string, increaseAmount int) (time.Time, error) {
	if err := validation.IsValidXID(userID); err != nil {
		s.logger.Error(entity.InvalidUserID, s.logger.StringModifier("userID", userID))
		return time.Time{}, err
	}
	if increaseAmount <= 0 || increaseAmount >= 10 {
		s.logger.Error("comment increase amount should be between 1-9 including 1 and 9")
		return time.Time{}, entity.IncreaseAmountNotValid
	}
	return s.repository.IncreaseUserPostCount(ctx, userID, increaseAmount)
}

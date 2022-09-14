package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/core/validation"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/rs/xid"
)

func (s *Service) Login(ctx context.Context, email string, username *string) (*entity.User, error) {
	if err := validation.IsValidEmail(email); err != nil {
		return nil, err
	}
	exists, err := s.repository.IsUserExistsByEmail(ctx, email)
	if err != nil {
		s.logger.Error("error is user exists by email from repository")
		return nil, err
	}
	if exists {
		s.logger.Debug("user getting from repository")
		return s.repository.GetUserByEmail(ctx, email)
	}
	if username == nil || validation.IsValidUsername(*username) != nil {
		s.logger.Error("username not found")
		return nil, entity.UserNotFound.Err()
	}
	exists, err = s.repository.IsUserExistsByUsername(ctx, *username)
	if err != nil {
		s.logger.Error("getting user exist by username from repository error")
		return nil, err
	}
	if exists {
		return nil, entity.UsernameAlreadyTaken.Err()
	}
	id := xid.New().String()
	//if validated then create the user
	createdAt, err := s.repository.CreateUser(ctx, id, email, *username)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return &entity.User{
		ID:        id,
		Email:     email,
		Username:  *username,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}, nil
}

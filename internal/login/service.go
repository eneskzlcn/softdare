package login

import (
	"context"
	"errors"
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/rs/xid"
	"time"
)

type LoginRepository interface {
	IsUserExistsByEmail(ctx context.Context, email string) (bool, error)
	IsUserExistsByUsername(ctx context.Context, username string) (bool, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, request CreateUserRequest) (time.Time, error)
}
type Service struct {
	repository LoginRepository
	logger     logger.Logger
}

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUsernameAlreadyTaken = errors.New("username already taken")
)

func NewService(logger logger.Logger, repository LoginRepository) *Service {
	if logger == nil {
		fmt.Println("Given logger to service is nil.")
		return nil
	}
	if repository == nil {
		logger.Error(ErrRepositoryNil)
		return nil
	}
	return &Service{repository: repository, logger: logger}
}

func (s *Service) Login(ctx context.Context, inp Input) (*entity.User, error) {
	exists, err := s.repository.IsUserExistsByEmail(ctx, inp.Email)
	if err != nil {
		s.logger.Error("error is user exists by email from repository")
		return nil, err
	}
	if exists {
		s.logger.Debug("user getting from repository")
		return s.repository.GetUserByEmail(ctx, inp.Email)
	}
	if inp.Username == nil {
		s.logger.Error("username not found")
		return nil, ErrUserNotFound
	}
	exists, err = s.repository.IsUserExistsByUsername(ctx, *inp.Username)
	if err != nil {
		s.logger.Error("getting user exist by username from repository error")
		return nil, err
	}
	if exists {
		return nil, ErrUsernameAlreadyTaken
	}
	// if both username and email is given bot not found in db then validate the username and email
	id := xid.New().String()
	request := CreateUserRequest{
		ID:       id,
		Email:    inp.Email,
		Username: *inp.Username,
	}
	if err = request.Validate(); err != nil {
		s.logger.Error(ErrValidation)
		return nil, ErrValidation
	}
	//if validated then create the user
	createdAt, err := s.repository.CreateUser(ctx, request)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return &entity.User{
		ID:        id,
		Email:     inp.Email,
		Username:  *inp.Username,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}, nil
}

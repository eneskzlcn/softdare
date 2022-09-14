package service

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/entity"
)

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return s.repository.GetUserByUsername(ctx, username)
}

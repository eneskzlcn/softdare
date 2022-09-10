package login_test

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/login"
	mocks "github.com/eneskzlcn/softdare/internal/mocks/login"
	"github.com/golang/mock/gomock"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestNewService(t *testing.T) {
	t.Run("given nil logger then it should return nil when NewService called", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockRepo := mocks.NewMockLoginRepository(controller)
		service := login.NewService(nil, mockRepo)
		assert.Nil(t, service)
	})
	t.Run("given nil repository then it should return nil when NewService called", func(t *testing.T) {
		service := login.NewService(zap.NewExample().Sugar(), nil)
		assert.Nil(t, service)
	})
	t.Run("given valid args then it should return new service when NewService called", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockRepo := mocks.NewMockLoginRepository(controller)
		service := login.NewService(zap.NewExample().Sugar(), mockRepo)
		assert.NotNil(t, service)
	})
}
func TestService_Login(t *testing.T) {
	controller := gomock.NewController(t)
	mockRepo := mocks.NewMockLoginRepository(controller)
	service := login.NewService(zap.NewExample().Sugar(), mockRepo)

	t.Run("given exist user with email then it should return the user without oops when Login called", func(t *testing.T) {
		inp := login.LoginInput{
			Email:    "found@email.com",
			Username: new(string),
		}
		ctx := context.Background()
		mockRepo.EXPECT().IsUserExistsByEmail(ctx, inp.Email).Return(true, nil)
		expectedUser := &login.User{
			ID:        xid.New().String(),
			Email:     inp.Email,
			Username:  *inp.Username,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mockRepo.EXPECT().GetUserByEmail(ctx, inp.Email).Return(expectedUser, nil)
		user, err := service.Login(ctx, inp)
		assert.Nil(t, err)
		assert.Equal(t, user, expectedUser)
	})
	t.Run("given not exist user with email and username is nil then should return nil and oops", func(t *testing.T) {
		inp := login.LoginInput{
			Email:    "found@email.com",
			Username: nil,
		}
		ctx := context.Background()
		mockRepo.EXPECT().IsUserExistsByEmail(ctx, inp.Email).Return(false, nil)
		user, err := service.Login(ctx, inp)
		assert.Nil(t, user)
		assert.Equal(t, err, login.ErrUserNotFound)
	})
	t.Run("given not exist user with email and given a username that already taken then should return nil and oops", func(t *testing.T) {
		username := "valid"
		inp := login.LoginInput{
			Email:    "found@email.com",
			Username: &username,
		}
		ctx := context.Background()
		mockRepo.EXPECT().IsUserExistsByEmail(ctx, inp.Email).Return(false, nil)
		mockRepo.EXPECT().IsUserExistsByUsername(ctx, username).Return(true, nil)
		user, err := service.Login(ctx, inp)
		assert.Nil(t, user)
		assert.Equal(t, err, login.ErrUsernameAlreadyTaken)
	})
	t.Run("given not exist email and username that not already taken and valid then should return new created user without oops", func(t *testing.T) {
		username := "valid"
		inp := login.LoginInput{
			Email:    "found@email.com",
			Username: &username,
		}
		createTime := time.Now()
		expectedUser := login.User{
			ID:        "",
			Email:     inp.Email,
			Username:  username,
			CreatedAt: createTime,
			UpdatedAt: createTime,
		}
		ctx := context.Background()
		mockRepo.EXPECT().IsUserExistsByEmail(ctx, inp.Email).Return(false, nil)
		mockRepo.EXPECT().IsUserExistsByUsername(ctx, username).Return(false, nil)
		mockRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(createTime, nil)
		user, err := service.Login(ctx, inp)
		assert.NotNil(t, user)
		assert.Nil(t, err)
		expectedUser.ID = user.ID
		assert.Equal(t, *user, expectedUser)
	})
	t.Run("given not exist email and username that not already taken but invalid then should return nil with oops", func(t *testing.T) {
		username := "no"
		inp := login.LoginInput{
			Email:    "notemail",
			Username: &username,
		}
		createTime := time.Now()
		ctx := context.Background()
		mockRepo.EXPECT().IsUserExistsByEmail(ctx, inp.Email).Return(false, nil)
		mockRepo.EXPECT().IsUserExistsByUsername(ctx, username).Return(false, nil)
		mockRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(createTime, nil)
		_, err := service.Login(ctx, inp)
		assert.NotNil(t, err)
	})
}

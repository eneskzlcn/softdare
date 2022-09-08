package post_test

import (
	"context"
	"fmt"
	mocks "github.com/eneskzlcn/softdare/internal/mocks/post"
	"github.com/eneskzlcn/softdare/internal/post"
	contextUtil "github.com/eneskzlcn/softdare/internal/util/context"
	"github.com/golang/mock/gomock"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"strings"
	"testing"
	"time"
)

func TestNewService(t *testing.T) {
	controller := gomock.NewController(t)
	mockPostRepo := mocks.NewMockPostRepository(controller)
	logger := zap.NewExample().Sugar()

	t.Run("given nil logger then it should return nil when NewService called", func(t *testing.T) {
		service := post.NewService(mockPostRepo, nil)
		assert.Nil(t, service)
	})
	t.Run("given nil repository then it should return nil when NewService called", func(t *testing.T) {
		service := post.NewService(nil, logger)
		assert.Nil(t, service)
	})
	t.Run("given valid args then it should return service when NewService called", func(t *testing.T) {
		service := post.NewService(mockPostRepo, logger)
		assert.NotNil(t, service)
	})
}
func TestService_CreatePost(t *testing.T) {
	controller := gomock.NewController(t)
	mockPostRepo := mocks.NewMockPostRepository(controller)
	logger := zap.NewExample().Sugar()
	service := post.NewService(mockPostRepo, logger)
	assert.NotNil(t, service)

	t.Run("given unauthorized (not exist) user then it should return nil and error when CreatePost called", func(t *testing.T) {
		validInput := post.CreatePostInput{
			Content: "I am a content",
		}
		resp, err := service.CreatePost(context.Background(), validInput)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
	t.Run("given authorized user with not valid post content then it should return nil and error when CreatePost called", func(t *testing.T) {
		notValidInput := post.CreatePostInput{
			Content: strings.Repeat("a", 1200),
		}
		ctx := contextUtil.WithContext[post.User](context.Background(), "user", post.User{
			ID:       xid.New().String(),
			Email:    "valid@gmail.com",
			Username: "validname",
		})
		resp, err := service.CreatePost(ctx, notValidInput)
		assert.NotNil(t, err)
		fmt.Println(err.Error())
		assert.Nil(t, resp)
	})
	t.Run("given authorized user with  valid post content then it should return created post without error when CreatePost called", func(t *testing.T) {
		validInput := post.CreatePostInput{
			Content: "I am a content",
		}
		createTime := time.Now()
		ctx := contextUtil.WithContext[post.User](context.Background(), "user", post.User{
			ID:       xid.New().String(),
			Email:    "valid@gmail.com",
			Username: "validname",
		})
		mockPostRepo.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Return(createTime, nil)
		resp, err := service.CreatePost(ctx, validInput)
		assert.Nil(t, err)
		assert.Equal(t, resp.CreatedAt, createTime)
	})
}

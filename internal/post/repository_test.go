package post_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eneskzlcn/softdare/internal/post"
	"github.com/eneskzlcn/softdare/postgres/mocks"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"regexp"
	"testing"
	"time"
)

func TestNewRepository(t *testing.T) {
	db, _ := mocks.NewMockPostgresDB()
	logger := zap.NewExample().Sugar()
	t.Run("given nil logger then it should return nil when NewRepository called", func(t *testing.T) {
		repo := post.NewRepository(db, nil)
		assert.Nil(t, repo)
	})
	t.Run("given nil db then it should return nil when NewRepository called", func(t *testing.T) {
		repo := post.NewRepository(nil, logger)
		assert.Nil(t, repo)
	})
	t.Run("given valid args then it should return repository when NewRepository called", func(t *testing.T) {
		repo := post.NewRepository(db, logger)
		assert.NotNil(t, repo)
	})
}
func TestRepository_CreatePost(t *testing.T) {
	db, mock := mocks.NewMockPostgresDB()
	logger := zap.NewExample().Sugar()
	repo := post.NewRepository(db, logger)
	assert.NotNil(t, repo)
	createPostQuery := `INSERT INTO posts (id, user_id, content) VALUES ($1, $2, $3) RETURNING created_at`
	postID := xid.New().String()
	userID := xid.New().String()
	createTime := time.Now()
	content := "It is my first post."
	mock.ExpectQuery(regexp.QuoteMeta(createPostQuery)).WithArgs(postID, userID, content).
		WillReturnRows(sqlmock.NewRows([]string{"created_at"}).AddRow(createTime))

	time, err := repo.CreatePost(context.Background(), post.CreatePostRequest{
		ID:      postID,
		UserID:  userID,
		Content: content,
	})
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
	assert.Equal(t, time, createTime)
}

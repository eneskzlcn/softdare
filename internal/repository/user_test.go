package repository_test

//
//import (
//	"context"
//	"errors"
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/eneskzlcn/softdare/internal/login"
//	"github.com/eneskzlcn/softdare/postgres/mocks"
//	"github.com/rs/xid"
//	"github.com/stretchr/testify/assert"
//	"go.uber.org/zap"
//	"regexp"
//	"testing"
//	"time"
//)
//
//func TestNewRepository(t *testing.T) {
//	t.Run("given nil logger then it should return nil", func(t *testing.T) {
//		db, _ := mocks.NewMockPostgresDB()
//		repo := login.NewRepository(nil, db)
//		assert.Nil(t, repo)
//	})
//	t.Run("given nil database then it should return nil", func(t *testing.T) {
//		repo := login.NewRepository(zap.NewExample().Sugar(), nil)
//		assert.Nil(t, repo)
//	})
//	t.Run("given valid args then it should return repository", func(t *testing.T) {
//		db, _ := mocks.NewMockPostgresDB()
//		repo := login.NewRepository(zap.NewExample().Sugar(), db)
//		assert.NotNil(t, repo)
//	})
//}
//func TestRepository_GetUserByEmail(t *testing.T) {
//	db, mock := mocks.NewMockPostgresDB()
//	repo := login.NewRepository(zap.NewExample().Sugar(), db)
//	assert.NotNil(t, repo)
//
//	const userByEmailQuery = `SELECT id, email, username, created_at, updated_at FROM users WHERE email = $1`
//
//	t.Run("given exist user email then it should return user without oops", func(t *testing.T) {
//		expectedUser := login.User{
//			ID:        xid.New().String(),
//			Email:     "test@gmail.com",
//			Username:  "username",
//			CreatedAt: time.Now(),
//			UpdatedAt: time.Now(),
//		}
//		mock.ExpectQuery(regexp.QuoteMeta(userByEmailQuery)).
//			WithArgs(expectedUser.Email).
//			WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "created_at", "updated_at"}).
//				AddRow(expectedUser.ID, expectedUser.Email, expectedUser.Username, expectedUser.CreatedAt, expectedUser.UpdatedAt))
//
//		user, err := repo.GetUserByEmail(context.Background(), expectedUser.Email)
//		assert.Nil(t, err)
//		err = mock.ExpectationsWereMet()
//		assert.Nil(t, err)
//		assert.Equal(t, *user, expectedUser)
//	})
//	t.Run("given not exist user email then it should return nil with oops", func(t *testing.T) {
//		expectedUser := login.User{
//			ID:        xid.New().String(),
//			Email:     "test@gmail.com",
//			Username:  "username",
//			CreatedAt: time.Now(),
//			UpdatedAt: time.Now(),
//		}
//		mock.ExpectQuery(regexp.QuoteMeta(userByEmailQuery)).
//			WithArgs(expectedUser.Email).
//			WillReturnError(errors.New("user not found"))
//
//		_, err := repo.GetUserByEmail(context.Background(), expectedUser.Email)
//		assert.NotNil(t, err)
//		err = mock.ExpectationsWereMet()
//		assert.Nil(t, err)
//	})
//}
//func TestRepository_IsUserExistsByEmail(t *testing.T) {
//	db, mock := mocks.NewMockPostgresDB()
//	repo := login.NewRepository(zap.NewExample().Sugar(), db)
//	assert.NotNil(t, repo)
//	email := ""
//	const userExistsByEmailQuery = `SELECT EXISTS ( SELECT 1 FROM users WHERE email = $1) `
//
//	t.Run("given exist user email then it should return true", func(t *testing.T) {
//		mock.ExpectQuery(regexp.QuoteMeta(userExistsByEmailQuery)).WithArgs(email).
//			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))
//
//		exists, err := repo.IsUserExistsByEmail(context.Background(), email)
//		assert.Nil(t, err)
//		assert.True(t, exists)
//	})
//	t.Run("given not exist user email then it should return false", func(t *testing.T) {
//		mock.ExpectQuery(regexp.QuoteMeta(userExistsByEmailQuery)).WithArgs(email).
//			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
//
//		exists, err := repo.IsUserExistsByEmail(context.Background(), email)
//		assert.Nil(t, err)
//		assert.False(t, exists)
//	})
//}
//func TestRepository_IsUserExistsByUsername(t *testing.T) {
//	db, mock := mocks.NewMockPostgresDB()
//	repo := login.NewRepository(zap.NewExample().Sugar(), db)
//	assert.NotNil(t, repo)
//	username := "ens"
//	const userExistsByUsernameQuery = `SELECT EXISTS ( SELECT 1 FROM users WHERE username ILIKE $1 )`
//
//	t.Run("given exist user name then it should return true", func(t *testing.T) {
//		mock.ExpectQuery(regexp.QuoteMeta(userExistsByUsernameQuery)).WithArgs(username).
//			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))
//
//		exists, err := repo.IsUserExistsByUsername(context.Background(), username)
//		assert.Nil(t, err)
//		assert.True(t, exists)
//	})
//	t.Run("given not exist user name then it should return false", func(t *testing.T) {
//		mock.ExpectQuery(regexp.QuoteMeta(userExistsByUsernameQuery)).WithArgs(username).
//			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))
//
//		exists, err := repo.IsUserExistsByUsername(context.Background(), username)
//		assert.Nil(t, err)
//		assert.False(t, exists)
//	})
//}
//func TestRepository_CreateUser(t *testing.T) {
//	db, mock := mocks.NewMockPostgresDB()
//	repo := login.NewRepository(zap.NewExample().Sugar(), db)
//	assert.NotNil(t, repo)
//	const createUserQuery = `INSERT INTO users (id, email, username) VALUES ($1, $2, $3) RETURNING created_at`
//	user := login.CreateUserRequest{
//		ID:       xid.New().String(),
//		Email:    "new@mail.com",
//		Username: "newuser",
//	}
//	createTime := time.Now()
//	mock.ExpectQuery(regexp.QuoteMeta(createUserQuery)).WithArgs(user.ID, user.Email, user.Username).
//		WillReturnRows(sqlmock.NewRows([]string{"created_at"}).AddRow(createTime))
//
//	createdAt, err := repo.CreateUser(context.Background(), user)
//	assert.Nil(t, err)
//	assert.Equal(t, createTime, createdAt)
//}

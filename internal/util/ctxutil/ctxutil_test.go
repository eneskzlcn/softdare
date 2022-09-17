package ctxutil_test

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/util/ctxutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromContext(t *testing.T) {
	type user struct {
		name string
		age  string
	}
	aUser := user{
		name: "enes",
		age:  "23",
	}
	ctx := context.WithValue(context.Background(), "user", aUser)
	foundUser, exists := ctxutil.FromContext[user]("user", ctx)
	assert.True(t, exists)
	assert.Equal(t, foundUser, aUser)
}

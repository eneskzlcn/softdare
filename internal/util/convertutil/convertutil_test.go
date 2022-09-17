package convertutil_test

import (
	"github.com/eneskzlcn/softdare/internal/util/convertutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyTo(t *testing.T) {
	type usr1 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	type usr2 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	user1 := usr1{
		Name: "enes",
		Age:  23,
	}
	user2, err := convertutil.AnyTo[usr2](user1)
	assert.Nil(t, err)
	assert.Equal(t, user2.Age, user1.Age)
	assert.Equal(t, user2.Name, user1.Name)
}

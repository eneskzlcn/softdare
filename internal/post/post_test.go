package post_test

import (
	"github.com/eneskzlcn/softdare/internal/post"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCreatePostInput_Validate(t *testing.T) {
	t.Run("given too long content then it should return oops when validate called", func(t *testing.T) {
		tooLongContentInp := post.CreatePostInput{Content: strings.Repeat("a", 1200)}
		err := tooLongContentInp.Validate()
		assert.NotNil(t, err)
	})
	t.Run("given too short content then it should return oops when validate called", func(t *testing.T) {
		tooShortContentInp := post.CreatePostInput{Content: "a"}
		err := tooShortContentInp.Validate()
		assert.NotNil(t, err)
	})
	t.Run("given valid content then it should return nil when validate called", func(t *testing.T) {
		validContentInp := post.CreatePostInput{Content: "I am a valid content"}
		err := validContentInp.Validate()
		assert.Nil(t, err)
	})
}
func TestCreatePostInput_Prepare(t *testing.T) {
	t.Run("given content has 2 consecutive spaces then it should remove one of them when Prepare called", func(t *testing.T) {
		inp := post.CreatePostInput{Content: "I am a  valid content"}
		inp.Prepare()
		assert.Equal(t, inp.Content, "I am a valid content")
	})
	t.Run("given content has unnecessary end of line then it should trim that spaces when Prepare called", func(t *testing.T) {
		inp := post.CreatePostInput{Content: "I am a valid content\n\n\n"}
		inp.Prepare()
		assert.Equal(t, inp.Content, "I am a valid content")
	})
}

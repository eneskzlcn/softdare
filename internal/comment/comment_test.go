package comment_test

import (
	"github.com/eneskzlcn/softdare/internal/comment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCommentInput_Validate(t *testing.T) {
	//type testCase struct {
	//}
	//testCases := []struct {
	//	scenario string
	//	content  string
	//	postID   string
	//	expected error
	//}{
	//	{scenario: "given extra spaces", ""},
	//}
}
func TestCreateCommentInput_Prepare(t *testing.T) {
	t.Run("given content that has more spaces and end of lines then it should remove all the unnecessary character when prepare called", func(t *testing.T) {
		inp := comment.CreateCommentInput{
			Content: "asd  asfsafaf \n\n\n adasdas d\n\n\n",
		}
		inp.Prepare()
		assert.Equal(t, "asd asfsafaf \n\n adasdas d", inp.Content)
	})
}

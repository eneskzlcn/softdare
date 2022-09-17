package input_test

import (
	"github.com/eneskzlcn/softdare/internal/core/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanUserContentInput(t *testing.T) {
	content := "Need to  get rid  of double  spaces and   \n double end of lines like\n\n and also unnecessary end of line like\n\n\n\n\n"
	res := input.CleanUserContentInput(content)
	expected := "Need to get rid of double spaces and  \n double end of lines like\n and also unnecessary end of line like"
	assert.Equal(t, expected, res)
}

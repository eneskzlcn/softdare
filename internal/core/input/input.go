package input

import (
	"strings"
)

func CleanUserContentInput(content string) string {
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "\n\n", "\n")
	return strings.ReplaceAll(content, "  ", " ")
}

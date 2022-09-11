package post

import (
	"fmt"
	"github.com/mvdan/xurls"
	"html/template"
	"time"
)

func FormatPostTime(createdAt time.Time) string {
	durationBetweenNow := time.Now().Sub(createdAt)
	durationMinutes := int(durationBetweenNow.Minutes())
	if durationMinutes < 60 && durationMinutes > 0 {
		return fmt.Sprintf("%dm ago", durationMinutes)
	} else if durationMinutes == 0 {
		return "Just Now"
	} else if durationMinutes >= 60 && durationMinutes < 1440 {
		return fmt.Sprintf("%dh ago", durationMinutes/60)
	} else if durationMinutes >= 1440 && durationMinutes < 10080 {
		return fmt.Sprintf("%dd ago", durationMinutes/1440)
	} else if durationMinutes >= 10080 && durationMinutes < 40320 {
		return fmt.Sprintf("%dw ago", durationMinutes/10080)
	} else if durationMinutes >= 40320 && durationMinutes < 483840 {
		return fmt.Sprintf("%dmonths ago", durationMinutes/40320)
	} else if durationMinutes >= 483840 {
		return fmt.Sprintf("%dyears ago", durationMinutes/483840)
	}
	return ""
}
func FormatPostContent(content string) template.HTML {
	content = template.HTMLEscapeString(content)
	return template.HTML(xurls.Relaxed.
		ReplaceAllString(content, `<a href ="$0" target="_blank" rel="noopener noreferror">$0</a>`))
}

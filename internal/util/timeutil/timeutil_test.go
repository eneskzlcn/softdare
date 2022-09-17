package timeutil_test

import (
	"github.com/eneskzlcn/softdare/internal/util/timeutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToAgoFormatter(t *testing.T) {
	t.Run("given time little than now but not more than 60 seconds then it should return Just Now", func(t *testing.T) {
		timeAgo := timeutil.ToAgoFormatter(time.Now())
		assert.Equal(t, timeAgo, "Just Now")
	})
	t.Run("given time little than now at least 60 seconds and at most 1 hour then it should return the time in xm ago where x is the minute", func(t *testing.T) {
		givenTime := time.Now().Add(-time.Minute * 2)
		timeAgo := timeutil.ToAgoFormatter(givenTime)
		assert.Equal(t, "2m ago", timeAgo)
	})
	t.Run("given time little than now at least 1 hour and at most 1 day then it should return the time in xh ago where x is the hour", func(t *testing.T) {
		givenTime := time.Now().Add(-time.Hour * 2)
		timeAgo := timeutil.ToAgoFormatter(givenTime)
		assert.Equal(t, "2h ago", timeAgo)
	})
	t.Run("given time little than now at least 1 day and at most 1 week then it should return the time in xd ago where d is the day", func(t *testing.T) {
		givenTime := time.Now().Add(-time.Hour * 27)
		timeAgo := timeutil.ToAgoFormatter(givenTime)
		assert.Equal(t, "1d ago", timeAgo)
	})
	t.Run("given time little than now at least 1 week and at most 1 month then it should return the time in xw ago where w is the week", func(t *testing.T) {
		givenTime := time.Now().Add(-time.Hour * 24 * 8)
		timeAgo := timeutil.ToAgoFormatter(givenTime)
		assert.Equal(t, "1w ago", timeAgo)
	})
	t.Run("given time little than now at least 1 month and at most 1 year then it should return the time in xmonths ago where w is the week", func(t *testing.T) {
		givenTime := time.Now().Add(-time.Hour * 24 * 32)
		timeAgo := timeutil.ToAgoFormatter(givenTime)
		assert.Equal(t, "1months ago", timeAgo)
	})
	t.Run("given time little than now at least 1 year and then it should return the time in xyears ago where x is the year", func(t *testing.T) {
		givenTime := time.Now().Add(-time.Hour * 24 * 30 * 13)
		timeAgo := timeutil.ToAgoFormatter(givenTime)
		assert.Equal(t, "1years ago", timeAgo)
	})
	t.Run("given time older than at least 1 year from now always be listed in xyears format where x is the year", func(t *testing.T) {
		givenTime := time.Now().Add(-time.Hour * 24 * 30 * 252)
		timeAgo := timeutil.ToAgoFormatter(givenTime)
		assert.Equal(t, "22years ago", timeAgo)
	})
}

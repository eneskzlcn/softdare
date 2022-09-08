package home_test

import (
	"github.com/eneskzlcn/softdare/internal/home"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionDataFromAny(t *testing.T) {
	sessionData := struct {
		Email    string
		Username string
	}{
		Email:    "me@gm.com",
		Username: "me",
	}
	userSessionData, err := home.SessionDataFromAny(sessionData)
	assert.Nil(t, err)
	assert.Equal(t, sessionData.Email, userSessionData.Email)
	assert.Equal(t, sessionData.Username, userSessionData.Username)
}

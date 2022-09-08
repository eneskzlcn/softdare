package config_test

import (
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	environment := "test"
	config, err := config.LoadConfig("../../.dev/", environment, "yaml")
	assert.Nil(t, err)
	assert.Equal(t, ":4000", config.Server.Address)
	assert.Equal(t, "softdare", config.Db.Username)
	assert.Equal(t, "u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4", config.Session.Key)
}

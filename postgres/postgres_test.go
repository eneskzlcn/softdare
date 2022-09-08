package postgres_test

import (
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/postgres"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGivenConfigThenItShouldCreatePostgresDatabaseConnectionSuccessfullyWhenNewCalled(t *testing.T) {
	//this test is changes for the environment that test is running on
	env := os.Getenv("DEPLOYMENT_ENVIRONMENT")
	envConfigMap := map[string]string{
		"":      "test",
		"local": "local",
		"prod":  "prod",
		"qa":    "qa",
	}
	configs, err := config.LoadConfig("../.dev/", envConfigMap[env], "yaml")
	assert.Nil(t, err)

	db, err := postgres.New(configs.Db)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

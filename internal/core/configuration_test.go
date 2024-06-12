package core

import (
	mim "github.com/ONSdigital/dp-mongodb-in-memory"
	"github.com/mblancoa/go-rpc/conf"
	"github.com/mblancoa/go-rpc/internal/errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var mongodbServer *mim.Server

func init() {
	err := os.Chdir("./../../")
	errors.ManageErrorPanic(err)
	err = os.Setenv(conf.RunMode, "test")
	errors.ManageErrorPanic(err)
}

func TestLoadConfiguration(t *testing.T) {
	var config coreConfiguration
	conf.LoadYamlConfiguration(conf.GetConfigFile(), &config)

	assert.NotEmpty(t, config)
	assert.NotEmpty(t, config.FileStorage)
	f := config.FileStorage
	assert.NotEmpty(t, f.StorageDire)
	assert.Equal(t, "directory", f.StorageDire)
	assert.NotEmpty(t, f.FilePattern)
	assert.Equal(t, "{{type}}-{{version}}.json", f.FilePattern)
}

func TestSetupCoreConfiguration(t *testing.T) {
	SetupCoreConfiguration()
	assert.NotEmpty(t, Context.InfoFileService)
}

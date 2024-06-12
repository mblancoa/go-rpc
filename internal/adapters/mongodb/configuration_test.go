package mongodb

import (
	"context"
	mim "github.com/ONSdigital/dp-mongodb-in-memory"
	"github.com/mblancoa/go-rpc/conf"
	"github.com/mblancoa/go-rpc/internal/core"
	"github.com/mblancoa/go-rpc/internal/errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var mongodbServer *mim.Server

func init() {
	err := os.Chdir("./../../..")
	errors.ManageErrorPanic(err)
	err = os.Setenv(conf.RunMode, "test")
	errors.ManageErrorPanic(err)
}

func setupDB() {
	server, err := mim.StartWithOptions(context.TODO(), "5.0.2", mim.WithPort(37017))
	errors.ManageErrorPanic(err)
	mongodbServer = server
}

func TearDownDB() {
	mongodbServer.Stop(context.TODO())
}

func TestLoadConfiguration(t *testing.T) {
	var config mongoDbConfiguration
	conf.LoadYamlConfiguration(conf.GetConfigFile(), &config)

	assert.NotEmpty(t, config)
	assert.NotEmpty(t, config.Mongodb)
	db := config.Mongodb.Database
	assert.NotEmpty(t, db)
	assert.Equal(t, "dbTest", db.Name)
	con := db.Connection
	assert.NotEmpty(t, con)
	assert.Equal(t, "localhost", con.Host)
	assert.Equal(t, int(37017), con.Port)
	assert.Equal(t, "mongodb", con.Username)
	assert.Equal(t, "TEST_DB_PASSWORD", con.Password)
}

func TestSetupMongodbConfiguration(t *testing.T) {
	setupDB()
	defer TearDownDB()

	SetupMongodbConfiguration()

	assert.NotEmpty(t, core.PersistenceContext)
	assert.NotEmpty(t, core.PersistenceContext.InfoFileRepository)
}

package mongodb

import (
	"context"
	"fmt"
	"github.com/mblancoa/go-rpc/conf"
	"github.com/mblancoa/go-rpc/internal/core"
	"github.com/mblancoa/go-rpc/internal/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type mongoDbConfiguration struct {
	Mongodb struct {
		Database struct {
			Name       string `yaml:"name"`
			Connection struct {
				Host     string `yaml:"host"`
				Port     int    `yaml:"port"`
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			} `yaml:"connection"`
		} `yaml:"database"`
	} `yaml:"mongodb"`
}

func SetupMongodbConfiguration() {
	log.Info().Msg("Initializing mongodb configuration")
	var config mongoDbConfiguration
	conf.LoadYamlConfiguration(conf.GetConfigFile(), &config)

	conn := config.Mongodb.Database.Connection
	connectionString := fmt.Sprintf("%s:%s//%s:%d", conn.Username, os.Getenv(conn.Password), conn.Host, conn.Port)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	errors.ManageErrorPanic(err)
	err = client.Ping(ctx, nil)
	errors.ManageErrorPanic(err)

	database := client.Database(config.Mongodb.Database.Name)
	setupPersistenceContext(database)
}

func setupPersistenceContext(database *mongo.Database) {
	mongoDbInfoFileRepository := NewMongoDbInfoFileRepository(database.Collection(InfoFileCollection))
	persistenceCtx := core.PersistenceContext
	persistenceCtx.InfoFileRepository = NewInfoFileRepository(mongoDbInfoFileRepository)
}

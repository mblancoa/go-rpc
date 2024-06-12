package mongodb

import (
	"context"
	mim "github.com/ONSdigital/dp-mongodb-in-memory"
	"github.com/mblancoa/go-rpc/internal/errors"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDBPersistenceSuite struct {
	suite.Suite
	server                    *mim.Server
	client                    *mongo.Client
	database                  *mongo.Database
	infoFileCollection        *mongo.Collection
	mongoDbInfoFileRepository MongoDbInfoFileRepository
}

func (suite *mongoDBPersistenceSuite) SetupSuite() {
	testCtx := context.Background()

	server, err := mim.Start(testCtx, "5.0.2")
	errors.ManageErrorPanic(err)
	suite.server = server

	client, err := mongo.Connect(testCtx, options.Client().ApplyURI(server.URI()))
	errors.ManageErrorPanic(err)
	//Use client as needed
	err = client.Ping(testCtx, nil)
	errors.ManageErrorPanic(err)
	suite.client = client
	suite.database = client.Database("store")
	suite.Assert()
}

func (suite *mongoDBPersistenceSuite) TearDownSuite() {
	ctx := context.TODO()
	defer suite.server.Stop(ctx)
	err := suite.client.Disconnect(ctx)
	errors.ManageErrorPanic(err)
}

func (suite *mongoDBPersistenceSuite) setupInfoFileCollection() {
	db := suite.database
	log.Debug().Msgf("Creating collection '%s'", InfoFileCollection)
	err := db.CreateCollection(context.TODO(), InfoFileCollection)
	errors.ManageErrorPanic(err)

	collection := db.Collection(InfoFileCollection)

	idIdx := []mongo.IndexModel{
		{
			Keys: bson.M{
				"_id": 1,
			},
		},
		{
			Keys: bson.D{
				{Key: "type", Value: 1},
				{Key: "version", Value: 1},
			}, Options: options.Index().SetUnique(true),
		},
	}
	s, err := collection.Indexes().CreateMany(context.TODO(), idIdx)
	errors.ManageErrorPanic(err)
	for _, str := range s {
		log.Debug().Msg(str)
	}

	suite.infoFileCollection = collection
	suite.mongoDbInfoFileRepository = NewMongoDbInfoFileRepository(suite.infoFileCollection)
}

func insertOne(coll *mongo.Collection, ctx context.Context, obj interface{}) {
	log.Debug().Msgf("Inserting %v", obj)
	_, err := coll.InsertOne(ctx, obj)
	errors.ManageErrorPanic(err)
}

func findOne(coll *mongo.Collection, ctx context.Context, property string, value, entity interface{}) {
	log.Debug().Msgf("Finding object from collection '%s'", coll.Name())
	err := coll.FindOne(ctx, bson.M{
		property: value,
	}, options.FindOne().SetSort(bson.M{})).Decode(entity)
	errors.ManageErrorPanic(err)
}

func findOneByFilter(coll *mongo.Collection, ctx context.Context, filter map[string]interface{}, entity interface{}) {
	log.Debug().Msgf("Finding object from collection '%s'", coll.Name())
	err := coll.FindOne(ctx,
		filter,
		options.FindOne().SetSort(bson.M{})).Decode(entity)
	errors.ManageErrorPanic(err)
}

func deleteAll(coll *mongo.Collection, ctx context.Context) {
	log.Debug().Msgf("Deleting all documents in collection '%s'", coll.Name())
	_, err := coll.DeleteMany(ctx, bson.D{})
	errors.ManageErrorPanic(err)
}

func count(coll *mongo.Collection, ctx context.Context) int64 {
	c, err := coll.CountDocuments(ctx, bson.D{})
	errors.ManageErrorPanic(err)
	return c
}

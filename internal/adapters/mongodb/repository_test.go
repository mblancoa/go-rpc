package mongodb

import (
	"context"
	"errors"
	"github.com/devfeel/mapper"
	"github.com/mblancoa/go-rpc/internal/core/domain"
	"github.com/mblancoa/go-rpc/internal/core/ports"
	errors2 "github.com/mblancoa/go-rpc/internal/errors"
	"github.com/pioz/faker"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

// CredentialsPersistenceService Tests

type infoFileRepositorySuite struct {
	mongoDBPersistenceSuite
	infoFileRepository ports.InfoFileRepository
}

func (suite *infoFileRepositorySuite) SetupSuite() {
	suite.mongoDBPersistenceSuite.SetupSuite()
	suite.setupInfoFileCollection()
	suite.infoFileRepository = NewInfoFileRepository(suite.mongoDbInfoFileRepository)
}

func (suite *infoFileRepositorySuite) SetupTest() {
	ctx := context.Background()
	deleteAll(suite.infoFileCollection, ctx)
}

func (suite *infoFileRepositorySuite) TearDownSuite() {
	suite.mongoDBPersistenceSuite.TearDownSuite()
}

func TestInfoFileRepositorySuite(t *testing.T) {
	suite.Run(t, new(infoFileRepositorySuite))
}

func (suite *infoFileRepositorySuite) TestSaveOrUpdateByTypeAndVersion_successfulWhenSave() {
	infoFile := &domain.InfoFile{}
	_ = faker.Build(infoFile)
	result, err := suite.infoFileRepository.SaveOrUpdateByTypeAndVersion(context.Background(), infoFile)

	suite.Assertions.NoError(err)
	suite.Assertions.True(result)

	insertedDB := &InfoFileDB{}
	filter := make(map[string]interface{})
	filter["type"] = infoFile.Type
	filter["version"] = infoFile.Version
	findOneByFilter(suite.infoFileCollection, context.Background(), filter, insertedDB)

	inserted := &domain.InfoFile{}
	_ = mapper.Mapper(insertedDB, inserted)
	suite.Assertions.Equal(infoFile, inserted)
}

func (suite *infoFileRepositorySuite) TestSaveOrUpdateByTypeAndVersion_successfulWhenUpdate() {
	infoFile := &domain.InfoFile{}
	object := map[string]interface{}{"arg1": "Lorem ipsum", "arg2": "Lorem ipsum dolor sit amet, consectetur"}
	infoFile.Content = map[string]interface{}(object)
	_ = faker.Build(infoFile)
	infoFileDB := &InfoFileDB{}
	_ = mapper.Mapper(infoFile, infoFileDB)
	insertOne(suite.infoFileCollection, context.Background(), infoFileDB)
	infoFile.Content = map[string]interface{}{"title": "Lorem ipsum",
		"text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
		"obj":  undefinedType{"pro1": "value1", "prop2": "value2"},
	}

	result, err := suite.infoFileRepository.SaveOrUpdateByTypeAndVersion(context.Background(), infoFile)

	suite.Assertions.NoError(err)
	suite.Assertions.True(result)

	updatedDB := &InfoFileDB{}
	filter := make(map[string]interface{})
	filter["type"] = infoFile.Type
	filter["version"] = infoFile.Version
	findOneByFilter(suite.infoFileCollection, context.Background(), filter, updatedDB)

	updated := &domain.InfoFile{}
	_ = mapper.Mapper(updatedDB, updated)
	suite.Assertions.Equal(infoFile, updated)
}

type infoFileRepositoryWithMockSuite struct {
	suite.Suite
	mongoDbInfoFileRepository *MockMongoDbInfoFileRepository
	infoFileRepository        ports.InfoFileRepository
}

func (suite *infoFileRepositoryWithMockSuite) SetupSuite() {
	suite.mongoDbInfoFileRepository = NewMockMongoDbInfoFileRepository(suite.T())
	suite.infoFileRepository = NewInfoFileRepository(suite.mongoDbInfoFileRepository)
}

func TestInfoFileRepositoryWithMockSuite(t *testing.T) {
	suite.Run(t, new(infoFileRepositoryWithMockSuite))
}

func (suite *infoFileRepositoryWithMockSuite) TestSaveOrUpdateByTypeAndVersion_failWhenErrorInserting() {
	ctx := context.Background()
	infoFile := &domain.InfoFile{}
	_ = faker.Build(infoFile)
	infoFileDB := &InfoFileDB{}
	_ = mapper.Mapper(infoFile, infoFileDB)
	internalError := errors.New("error inserting")
	suite.mongoDbInfoFileRepository.EXPECT().FindByTypeAndVersion(ctx, infoFile.Type, infoFile.Version).Return(nil, mongo.ErrNoDocuments)
	suite.mongoDbInfoFileRepository.EXPECT().InsertOne(ctx, infoFileDB).Return(nil, internalError)

	result, err := suite.infoFileRepository.SaveOrUpdateByTypeAndVersion(ctx, infoFile)

	suite.Assertions.False(result)
	suite.Assertions.Error(err)
	suite.Assertions.Equal(errors2.ErrGenericError, err)
}

func (suite *infoFileRepositoryWithMockSuite) TestSaveOrUpdateByTypeAndVersion_failWhenErrorFinding() {
	ctx := context.Background()
	infoFile := &domain.InfoFile{}
	_ = faker.Build(infoFile)
	internalError := errors.New("error finding")
	suite.mongoDbInfoFileRepository.EXPECT().FindByTypeAndVersion(ctx, infoFile.Type, infoFile.Version).Return(nil, internalError)

	result, err := suite.infoFileRepository.SaveOrUpdateByTypeAndVersion(ctx, infoFile)

	suite.Assertions.False(result)
	suite.Assertions.Error(err)
	suite.Assertions.Equal(errors2.ErrGenericError, err)
}

func (suite *infoFileRepositoryWithMockSuite) TestSaveOrUpdateByTypeAndVersion_failWhenErrorUpdating() {
	ctx := context.Background()
	infoFile := &domain.InfoFile{}
	_ = faker.Build(infoFile)
	infoFileDB := &InfoFileDB{}
	_ = mapper.Mapper(infoFile, infoFileDB)
	internalError := errors.New("error updating")
	suite.mongoDbInfoFileRepository.EXPECT().FindByTypeAndVersion(ctx, infoFile.Type, infoFile.Version).Return(infoFileDB, nil)
	suite.mongoDbInfoFileRepository.EXPECT().UpdateById(ctx, infoFileDB, infoFileDB.Id).Return(false, internalError)

	result, err := suite.infoFileRepository.SaveOrUpdateByTypeAndVersion(ctx, infoFile)

	suite.Assertions.False(result)
	suite.Assertions.Error(err)
	suite.Assertions.Equal(errors2.ErrGenericError, err)
}

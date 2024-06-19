package rpc

import (
	"context"
	"github.com/mblancoa/go-rpc/internal/core"
	"github.com/mblancoa/go-rpc/internal/core/domain"
	errors2 "github.com/mblancoa/go-rpc/internal/errors"
	"github.com/pioz/faker"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"testing"
)

type infoFileServiceServerSuite struct {
	suite.Suite
	fileDir, filePattern string
	infoFileService      *core.MockInfoFileService
	serviceServer        InfoFileServiceServer
}

func (suite *infoFileServiceServerSuite) SetupSuite() {
	suite.infoFileService = core.NewMockInfoFileService(suite.T())
	suite.serviceServer = NewInfoFileServiceServer(suite.infoFileService)
}

func TestInfoFileServiceSuite(t *testing.T) {
	suite.Run(t, new(infoFileServiceServerSuite))
}

func (suite *infoFileServiceServerSuite) TestLoadFile_successful() {
	request := &InfoFileRequest{}
	_ = faker.Build(request)
	infoFile := &domain.InfoFile{
		Type:    request.Type,
		Version: request.Version,
		Hash:    request.Hash,
	}
	infoFileRS := &domain.InfoFile{
		Type:    request.Type,
		Version: request.Version,
		Hash:    request.Hash,
		Content: map[string]interface{}{"compilerOptions": map[string]interface{}{
			"target":                           "es2016",
			"module":                           "commonjs",
			"esModuleInterop":                  true,
			"forceConsistentCasingInFileNames": true,
			"strict":                           true,
			"skipLibCheck":                     true,
		}},
	}

	suite.infoFileService.EXPECT().LoadFile(infoFile).Times(1).Return(infoFileRS, nil)

	result, err := suite.serviceServer.LoadFile(context.Background(), request)

	suite.Assertions.NoError(err)
	suite.Assertions.NotEmpty(result)

	suite.Assertions.Equal(infoFileRS.Type, result.Type)
	suite.Assertions.Equal(infoFileRS.Version, result.Version)
	suite.Assertions.Equal(infoFileRS.Hash, result.Hash)
	suite.Assertions.NotEmpty(result.Content)
	content := result.Content.AsMap()
	suite.Assertions.Equal(infoFileRS.Content, content)
}

func (suite *infoFileServiceServerSuite) TestLoadFile_successfulWhenHashesAreDifferent() {
	request := &InfoFileRequest{}
	_ = faker.Build(request)
	infoFile := &domain.InfoFile{
		Type:    request.Type,
		Version: request.Version,
		Hash:    request.Hash,
	}
	infoFileRS := &domain.InfoFile{
		Type:    request.Type,
		Version: request.Version,
		Hash:    "7b0a202022636f6d70696c65724f7074696f6e73223a207b0a2020202022746172676574223a2022657332303136222c0a20202020226d6f64756c65223a2022636f6d6d6f6e6a73222c0a202020202265734d6f64756c65496e7465726f70223a20747275652c0a2020202022666f726365436f6e73697374656e74436173696e67496e46696c654e616d6573223a20747275652c0a2020202022737472696374223a20747275652c0a2020202022736b69704c6962436865636b223a20747275650a20207d0a7d0ae3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		Content: map[string]interface{}{"compilerOptions": map[string]interface{}{
			"target":                           "es2016",
			"module":                           "commonjs",
			"esModuleInterop":                  true,
			"forceConsistentCasingInFileNames": true,
			"strict":                           true,
			"skipLibCheck":                     true,
		}},
	}

	suite.infoFileService.EXPECT().LoadFile(infoFile).Times(1).Return(infoFileRS, nil)

	result, err := suite.serviceServer.LoadFile(context.Background(), request)

	suite.Assertions.NoError(err)
	suite.Assertions.NotEmpty(result)

	suite.Assertions.Equal(infoFileRS.Type, result.Type)
	suite.Assertions.Equal(infoFileRS.Version, result.Version)
	suite.Assertions.Equal(infoFileRS.Hash, result.Hash)
	suite.Assertions.Empty(result.Content)
}

func (suite *infoFileServiceServerSuite) TestLoadFile_successfulWhenTypeIsEmpty() {
	request := &InfoFileRequest{}
	_ = faker.Build(request)
	request.Type = ""
	infoFile := &domain.InfoFile{
		Type:    "core",
		Version: request.Version,
		Hash:    request.Hash,
	}
	infoFileRS := &domain.InfoFile{
		Type:    "core",
		Version: request.Version,
		Hash:    "7b0a202022636f6d70696c65724f7074696f6e73223a207b0a2020202022746172676574223a2022657332303136222c0a20202020226d6f64756c65223a2022636f6d6d6f6e6a73222c0a202020202265734d6f64756c65496e7465726f70223a20747275652c0a2020202022666f726365436f6e73697374656e74436173696e67496e46696c654e616d6573223a20747275652c0a2020202022737472696374223a20747275652c0a2020202022736b69704c6962436865636b223a20747275650a20207d0a7d0ae3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		Content: map[string]interface{}{"compilerOptions": map[string]interface{}{
			"target":                           "es2016",
			"module":                           "commonjs",
			"esModuleInterop":                  true,
			"forceConsistentCasingInFileNames": true,
			"strict":                           true,
			"skipLibCheck":                     true,
		}},
	}

	suite.infoFileService.EXPECT().LoadFile(infoFile).Times(1).Return(infoFileRS, nil)

	result, err := suite.serviceServer.LoadFile(context.Background(), request)

	suite.Assertions.NoError(err)
	suite.Assertions.NotEmpty(result)

	suite.Assertions.Equal(infoFileRS.Type, result.Type)
	suite.Assertions.Equal(infoFileRS.Version, result.Version)
	suite.Assertions.Equal(infoFileRS.Hash, result.Hash)
	suite.Assertions.Empty(result.Content)
}

func (suite *infoFileServiceServerSuite) TestLoadFile_successfulWhenVersionIsEmpty() {
	request := &InfoFileRequest{}
	_ = faker.Build(request)
	request.Version = ""
	infoFile := &domain.InfoFile{
		Type:    request.Type,
		Version: "1.0.0",
		Hash:    request.Hash,
	}
	infoFileRS := &domain.InfoFile{
		Type:    request.Type,
		Version: "1.0.0",
		Hash:    "7b0a202022636f6d70696c65724f7074696f6e73223a207b0a2020202022746172676574223a2022657332303136222c0a20202020226d6f64756c65223a2022636f6d6d6f6e6a73222c0a202020202265734d6f64756c65496e7465726f70223a20747275652c0a2020202022666f726365436f6e73697374656e74436173696e67496e46696c654e616d6573223a20747275652c0a2020202022737472696374223a20747275652c0a2020202022736b69704c6962436865636b223a20747275650a20207d0a7d0ae3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		Content: map[string]interface{}{"compilerOptions": map[string]interface{}{
			"target":                           "es2016",
			"module":                           "commonjs",
			"esModuleInterop":                  true,
			"forceConsistentCasingInFileNames": true,
			"strict":                           true,
			"skipLibCheck":                     true,
		}},
	}

	suite.infoFileService.EXPECT().LoadFile(infoFile).Times(1).Return(infoFileRS, nil)

	result, err := suite.serviceServer.LoadFile(context.Background(), request)

	suite.Assertions.NoError(err)
	suite.Assertions.NotEmpty(result)

	suite.Assertions.Equal(infoFileRS.Type, result.Type)
	suite.Assertions.Equal(infoFileRS.Version, result.Version)
	suite.Assertions.Equal(infoFileRS.Hash, result.Hash)
	suite.Assertions.Empty(result.Content)
}

func (suite *infoFileServiceServerSuite) TestLoadFile_failWhenServiceReturnsNotfoundError() {
	request := &InfoFileRequest{}
	_ = faker.Build(request)
	infoFile := &domain.InfoFile{
		Type:    request.Type,
		Version: request.Version,
		Hash:    request.Hash,
	}

	suite.infoFileService.EXPECT().LoadFile(infoFile).Times(1).Return(&domain.InfoFile{}, errors2.ErrNotFound)

	result, err := suite.serviceServer.LoadFile(context.Background(), request)

	suite.Assertions.Error(err)
	suite.Assertions.Empty(result)

	expectedError := NewRpcError(codes.NotFound, "not found")
	suite.Assertions.Equal(expectedError, err)
}

func (suite *infoFileServiceServerSuite) TestLoadFile_failWhenServiceReturnsGenericError() {
	request := &InfoFileRequest{}
	_ = faker.Build(request)
	infoFile := &domain.InfoFile{
		Type:    request.Type,
		Version: request.Version,
		Hash:    request.Hash,
	}

	suite.infoFileService.EXPECT().LoadFile(infoFile).Times(1).Return(&domain.InfoFile{}, errors2.ErrGenericError)

	result, err := suite.serviceServer.LoadFile(context.Background(), request)

	suite.Assertions.Error(err)
	suite.Assertions.Empty(result)

	expectedError := NewRpcError(codes.Unknown, "unexpected error")
	suite.Assertions.Equal(expectedError, err)
}

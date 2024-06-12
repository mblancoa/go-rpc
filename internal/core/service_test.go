package core

import (
	errors2 "errors"
	"github.com/mblancoa/go-rpc/internal/core/domain"
	"github.com/mblancoa/go-rpc/internal/core/ports"
	"github.com/mblancoa/go-rpc/internal/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type infoFileServiceSuite struct {
	suite.Suite
	fileDir, filePattern string
	infoFileRepository   *ports.MockInfoFileRepository
	infoFileService      InfoFileService
}

func (suite *infoFileServiceSuite) SetupSuite() {
	suite.fileDir = "./resources"
	suite.filePattern = "{{type}}-{{version}}.json"
	suite.infoFileRepository = ports.NewMockInfoFileRepository(suite.T())
	suite.infoFileService = NewInfoFileService(suite.fileDir, suite.filePattern, suite.infoFileRepository)
}

func TestInfoFileServiceSuite(t *testing.T) {
	suite.Run(t, new(infoFileServiceSuite))
}

func (suite *infoFileServiceSuite) TestLoadFile_successful() {
	infoFile := &domain.InfoFile{
		Type:    "core",
		Version: "1.0.0",
	}
	fullInfoFile := &domain.InfoFile{
		Type:    "core",
		Version: "1.0.0",
		Content: map[string]interface{}{"compilerOptions": map[string]interface{}{
			"target":                           "es2016",
			"module":                           "commonjs",
			"esModuleInterop":                  true,
			"forceConsistentCasingInFileNames": true,
			"strict":                           true,
			"skipLibCheck":                     true,
		}},
		Hash: "7b0a202022636f6d70696c65724f7074696f6e73223a207b0a2020202022746172676574223a2022657332303136222c0a20202020226d6f64756c65223a2022636f6d6d6f6e6a73222c0a202020202265734d6f64756c65496e7465726f70223a20747275652c0a2020202022666f726365436f6e73697374656e74436173696e67496e46696c654e616d6573223a20747275652c0a2020202022737472696374223a20747275652c0a2020202022736b69704c6962436865636b223a20747275650a20207d0a7d0ae3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	}

	suite.infoFileRepository.EXPECT().SaveOrUpdateByTypeAndVersion(mock.Anything, fullInfoFile).Times(1).Return(true, nil)

	result, err := suite.infoFileService.LoadFile(infoFile)

	suite.Assertions.NoError(err)
	suite.Assertions.NotEmpty(result)
}

func (suite *infoFileServiceSuite) TestLoadFile_failWhenJsonFileIsMalformed() {
	infoFile := &domain.InfoFile{
		Type:    "badjson",
		Version: "1.0.0",
	}

	result, err := suite.infoFileService.LoadFile(infoFile)

	suite.Assertions.Empty(result)
	suite.Assertions.Error(err)
	suite.Assertions.Equal(errors.ErrGenericError, err)
}

func (suite *infoFileServiceSuite) TestLoadFile_failWhenFileNotFound() {
	infoFile := &domain.InfoFile{
		Type:    "notfound",
		Version: "1.0.0",
	}

	result, err := suite.infoFileService.LoadFile(infoFile)

	suite.Assertions.Empty(result)
	suite.Assertions.Error(err)
	suite.Assertions.Equal(errors.ErrNotFound, err)
}

func (suite *infoFileServiceSuite) TestLoadFile_failWhenSaveOrUpdateByTypeAndVersionFails() {
	infoFile := &domain.InfoFile{
		Type:    "core",
		Version: "1.0.0",
	}
	fullInfoFile := &domain.InfoFile{
		Type:    "core",
		Version: "1.0.0",
		Content: map[string]interface{}{"compilerOptions": map[string]interface{}{
			"target":                           "es2016",
			"module":                           "commonjs",
			"esModuleInterop":                  true,
			"forceConsistentCasingInFileNames": true,
			"strict":                           true,
			"skipLibCheck":                     true,
		}},
		Hash: "7b0a202022636f6d70696c65724f7074696f6e73223a207b0a2020202022746172676574223a2022657332303136222c0a20202020226d6f64756c65223a2022636f6d6d6f6e6a73222c0a202020202265734d6f64756c65496e7465726f70223a20747275652c0a2020202022666f726365436f6e73697374656e74436173696e67496e46696c654e616d6573223a20747275652c0a2020202022737472696374223a20747275652c0a2020202022736b69704c6962436865636b223a20747275650a20207d0a7d0ae3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	}
	internalError := errors2.New("unexpected error")
	suite.infoFileRepository.EXPECT().SaveOrUpdateByTypeAndVersion(mock.Anything, fullInfoFile).Times(1).Return(false, internalError)

	result, err := suite.infoFileService.LoadFile(infoFile)

	suite.Assertions.Empty(result)
	suite.Assertions.Error(err)
	suite.Assertions.Equal(errors.ErrGenericError, internalError)
}

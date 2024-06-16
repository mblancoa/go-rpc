package rpc

import (
	"context"
	"github.com/mblancoa/go-rpc/internal/core"
	"github.com/mblancoa/go-rpc/internal/core/domain"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	defaultType    = "core"
	defaultVersion = "1.0.0"
)

type infoFileServiceServer struct {
	UnimplementedInfoFileServiceServer
	infoFileService core.InfoFileService
}

func NewInfoFileServiceServer(infoFileService core.InfoFileService) InfoFileServiceServer {
	return &infoFileServiceServer{
		infoFileService: infoFileService,
	}
}

func (s *infoFileServiceServer) LoadFile(ctx context.Context, request *InfoFileRequest) (*InfoFileResponse, error) {

	infoFile := &domain.InfoFile{
		Type:    getStringOrDefault(request.GetType(), defaultType),
		Version: getStringOrDefault(request.GetVersion(), defaultVersion),
		Hash:    request.GetHash(),
	}
	result, err := s.infoFileService.LoadFile(infoFile)
	if err != nil {
		return &InfoFileResponse{}, err
	}

	response := &InfoFileResponse{
		Type:    result.Type,
		Version: result.Version,
		Hash:    result.Hash,
	}
	if request.GetHash() == result.Hash {
		content, _ := structpb.NewStruct(result.Content)
		response.Content = content
	}

	return response, nil
}

func getStringOrDefault(value, def string) string {
	if value == "" {
		return def
	}
	return value
}

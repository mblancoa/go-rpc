package rpc

import (
	"context"
	"errors"
	"github.com/mblancoa/go-rpc/internal/core"
	"github.com/mblancoa/go-rpc/internal/core/domain"
	errors2 "github.com/mblancoa/go-rpc/internal/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	defaultType    = "core"
	defaultVersion = "1.0.0"
)

type rpcError struct {
	status *status.Status
}

func (e *rpcError) Error() string {
	return e.status.Message()
}
func (e *rpcError) GRPCStatus() *status.Status {
	return e.status
}
func NewRpcError(c codes.Code, msg string) error {
	return &rpcError{
		status: status.New(c, msg),
	}
}

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
		response := &InfoFileResponse{}
		if errors.Is(err, errors2.ErrNotFound) {
			return response, NewRpcError(codes.NotFound, err.Error())
		} else {
			return response, NewRpcError(codes.Unknown, err.Error())
		}
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

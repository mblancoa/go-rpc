package core

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/devfeel/mapper"
	"github.com/mblancoa/go-rpc/internal/core/domain"
	"github.com/mblancoa/go-rpc/internal/core/ports"
	errors2 "github.com/mblancoa/go-rpc/internal/errors"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"strings"
)

type InfoFileService interface {
	// LoadFile loads the file defined in infoFile and persists its content
	LoadFile(infoFile *domain.InfoFile) (*domain.InfoFile, error)
}

type infoFileService struct {
	fileDir, filePattern string
	infoFileRepository   ports.InfoFileRepository
}

// NewInfoFileService creates a new InfoFileService
// filePattern has to be similar to '{{type}}-{{version}}.json'
func NewInfoFileService(fileDir, filePattern string, infoFileRepository ports.InfoFileRepository) InfoFileService {
	pattern := strings.ReplaceAll(filePattern, "{{type}}", "%1s")
	pattern = strings.ReplaceAll(pattern, "{{version}}", "%2s")
	service := infoFileService{
		fileDir:            fileDir,
		filePattern:        pattern,
		infoFileRepository: infoFileRepository,
	}
	return &service
}

func (s *infoFileService) LoadFile(infoFile *domain.InfoFile) (*domain.InfoFile, error) {
	fileName := fmt.Sprintf(s.filePattern, infoFile.Type, infoFile.Version)
	file, err := openFile(s.fileDir, fileName)
	if err != nil {
		return &domain.InfoFile{}, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()

	result := &domain.InfoFile{}
	_ = mapper.Mapper(infoFile, result)

	content, err := io.ReadAll(file)
	if err != nil {
		log.Error().Msgf("unexpected error reading %s: %s", fileName, err)
		return &domain.InfoFile{}, errors2.ErrGenericError
	}

	result.Hash = getHash(content)
	err = json.Unmarshal(content, &result.Content)
	if err != nil {
		log.Error().Msgf("unexpected error Unmarshaling %s: %s", fileName, err)
		return &domain.InfoFile{}, errors2.ErrGenericError
	}

	_, err = s.infoFileRepository.SaveOrUpdateByTypeAndVersion(context.Background(), result)
	if err != nil {
		return &domain.InfoFile{}, err
	}

	return result, nil
}

func openFile(fileDir, fileName string) (*os.File, error) {
	filePath := fmt.Sprintf("%s/%s", fileDir, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Error().Msgf("file %s not found", filePath)
			return nil, errors2.ErrNotFound
		} else {
			log.Error().Msgf("unexpected error loading file %s: %s", filePath, err)
			return nil, errors2.ErrGenericError
		}
	}
	return file, nil
}

func getHash(bts []byte) string {
	h := sha256.New()
	return fmt.Sprintf("%x", h.Sum(bts))
}

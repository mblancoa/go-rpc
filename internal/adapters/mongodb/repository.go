package mongodb

import (
	"context"
	"errors"
	"github.com/devfeel/mapper"
	"github.com/mblancoa/go-rpc/internal/core/domain"
	"github.com/mblancoa/go-rpc/internal/core/ports"
	errors2 "github.com/mblancoa/go-rpc/internal/errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type infoFileRepository struct {
	mongoDbInfoFileRepository MongoDbInfoFileRepository
}

func NewInfoFileRepository(mongoDbInfoFileRepository MongoDbInfoFileRepository) ports.InfoFileRepository {
	return &infoFileRepository{
		mongoDbInfoFileRepository: mongoDbInfoFileRepository,
	}
}
func (r *infoFileRepository) SaveOrUpdateByTypeAndVersion(ctx context.Context, infoFile *domain.InfoFile) (bool, error) {
	result := true
	infoFileDB, err := r.mongoDbInfoFileRepository.FindByTypeAndVersion(ctx, infoFile.Type, infoFile.Version)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			toInsert := &InfoFileDB{}
			_ = mapper.Mapper(infoFile, toInsert)
			_, err = r.mongoDbInfoFileRepository.InsertOne(ctx, toInsert)
			if err != nil {
				log.Error().Msgf("Error inserting InfoFile object: %s", err)
				return false, errors2.ErrGenericError
			}
		} else {
			log.Error().Msgf("Error finding InfoFile object: %s", err)
			return false, errors2.ErrGenericError
		}
	} else {
		_ = mapper.Mapper(infoFile, infoFileDB)

		result, err = r.mongoDbInfoFileRepository.UpdateById(ctx, infoFileDB, infoFileDB.Id)
		if err != nil {
			log.Error().Msgf("Error updating InfoFile object: %s", err)
			return false, errors2.ErrGenericError
		}
	}
	return result, nil
}

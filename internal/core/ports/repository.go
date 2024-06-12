package ports

import (
	"context"
	"github.com/mblancoa/go-rpc/internal/core/domain"
)

type InfoFileRepository interface {
	//FindByTypeAndVersion(ctx context.Context, tp, version string) (*domain.InfoFile, error)
	//FindByHash(ctx context.Context, hash string) ([]*domain.InfoFile, error)
	SaveOrUpdateByTypeAndVersion(ctx context.Context, file *domain.InfoFile) (bool, error)
}

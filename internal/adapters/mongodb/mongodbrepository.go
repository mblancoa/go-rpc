package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	InfoFileCollection string = "infoFile"
)

type undefinedType map[string]interface{}

type InfoFileDB struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Type    string             `bson:"type"`
	Version string             `bson:"version"`
	Hash    string             `bson:"hash"`
	Content undefinedType      `bson:"content"`
}

func (i *InfoFileDB) String() string {
	return fmt.Sprintf("InfoFileDB:{Id: %s, Type: %s, Version: %s, Hash: %s, Content: %v",
		i.Id, i.Type, i.Version, i.Hash, i.Content)
}

//go:generate repogen -dest=mongodbinfofilerepository_impl.go -model=InfoFileDB -repo=MongoDbInfoFileRepository
type MongoDbInfoFileRepository interface {
	InsertOne(ctx context.Context, file *InfoFileDB) (interface{}, error)
	FindByTypeAndVersion(ctx context.Context, tp, version string) (*InfoFileDB, error)
	FindByHash(ctx context.Context, hash string) ([]*InfoFileDB, error)
	UpdateById(ctx context.Context, file *InfoFileDB, id primitive.ObjectID) (bool, error)
}

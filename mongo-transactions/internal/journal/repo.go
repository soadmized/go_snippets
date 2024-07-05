package journal

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoMongo struct {
	coll *mongo.Collection
}

func NewRepo(coll *mongo.Collection) *RepoMongo {
	return &RepoMongo{coll: coll}
}

func (r *RepoMongo) Insert(ctx context.Context, rec Record) error {
	_, err := r.coll.InsertOne(ctx, rec)
	if err != nil {
		return errors.Wrap(err, "insert log record")
	}

	return nil
}

package user

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewRepo(coll *mongo.Collection) *RepoMongo {
	return &RepoMongo{coll: coll}
}

type RepoMongo struct {
	coll *mongo.Collection
}

func (r *RepoMongo) Get(ctx context.Context, id int32) (*User, error) {
	var user User

	res := r.coll.FindOne(ctx, bson.M{"_id": bson.M{"$eq": id}})

	err := res.Decode(&user)
	if err != nil {
		return nil, errors.Wrap(err, "decode user to model")
	}

	return &user, nil
}

func (r *RepoMongo) Upsert(ctx context.Context, user User) error {
	filter := bson.M{"_id": bson.M{"$eq": user.ID}}
	upd := bson.M{"$set": bson.M{"name": user.Name, "comment": user.Comment}}

	_, err := r.coll.UpdateOne(ctx, filter, upd, options.Update().SetUpsert(true))
	if err != nil {
		return errors.Wrap(err, "upsert user")
	}

	return nil
}

func (r *RepoMongo) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	session, err := r.coll.Database().Client().StartSession(options.Session())
	if err != nil {
		return errors.Wrap(err, "start session")
	}

	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(ctxSession mongo.SessionContext) (interface{}, error) {
		return nil, fn(ctxSession)
	}, options.Transaction().SetReadPreference(readpref.Primary()))
	if err != nil {
		return errors.Wrap(err, "exec transaction")
	}

	return nil
}

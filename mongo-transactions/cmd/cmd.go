package cmd

import (
	"context"
	"fmt"
	"math/rand"
	"mongo-transactions/internal/user"

	"mongo-transactions/internal/journal"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURI    = "mongodb://localhost:27017"
	database    = "transactions"
	collUser    = "user"
	collJournal = "journal"
)

func Run(ctx context.Context) error {
	db, err := mongoDB(ctx, mongoURI, database)
	if err != nil {
		return errors.Wrap(err, "get mongo db")
	}

	srv, err := buildSrv(db)
	if err != nil {
		return errors.Wrap(err, "build service")
	}

	// Test run
	id1 := rand.Int31()
	user1 := user.User{
		ID:      id1,
		Name:    fmt.Sprintf("Travis Bickle - %d", id1),
		Comment: fmt.Sprintf("Travis Bickle - %d", id1),
	}

	id2 := rand.Int31()
	user2 := user.User{
		ID:      id2,
		Name:    fmt.Sprintf("Tony Montana - %d", id2),
		Comment: fmt.Sprintf("Tony Montana - %d", id2),
	}

	if err = srv.Upsert(ctx, user1); err != nil {
		return errors.Wrap(err, "upsert user1")
	}

	if err = srv.Upsert(ctx, user2); err != nil {
		return errors.Wrap(err, "upsert user2")
	}

	if err = db.Drop(ctx); err != nil {
		return errors.Wrap(err, "drop test db")
	}

	return nil
}

func buildSrv(db *mongo.Database) (*user.Service, error) {
	userColl, err := mongoColl(db, collUser)
	if err != nil {
		return nil, errors.Wrap(err, "get user coll")
	}

	journalColl, err := mongoColl(db, collJournal)
	if err != nil {
		return nil, errors.Wrap(err, "get journal coll")
	}

	journalSrv := journalService(journalColl)

	return userService(userColl, journalSrv), nil
}

func mongoColl(db *mongo.Database, collName string) (*mongo.Collection, error) {
	coll := db.Collection(collName)

	return coll, nil
}

func mongoDB(ctx context.Context, uri, dbName string) (*mongo.Database, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "connect to mongo")
	}

	db := conn.Database(dbName)

	return db, nil
}

func journalService(coll *mongo.Collection) *journal.Service {
	repo := journal.NewRepo(coll)

	return journal.NewService(repo)
}

func userService(coll *mongo.Collection, journalSrv *journal.Service) *user.Service {
	repo := user.NewRepo(coll)

	return user.NewService(repo, journalSrv)
}

package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepoTestSuite struct {
	suite.Suite
	conn *mongo.Client
	db   *mongo.Database
	coll *mongo.Collection
}

func (s *RepoTestSuite) SetupSuite() {
	const mongoURI = "mongodb://localhost:27017"

	ctx := context.Background()

	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	s.Require().NoError(err)
	s.conn = conn

	db := conn.Database(fmt.Sprintf("transactions-%s", uuid.New().String()))
	s.db = db

	coll := db.Collection("users-test")
	s.coll = coll

	// fixtures
	_, err = s.coll.InsertOne(ctx, bson.M{"_id": 1, "name": "John Snow", "comment": "snow warrior"})
	s.Require().NoError(err)

	_, err = s.coll.InsertOne(ctx, bson.M{"_id": 2, "name": "john_snow", "comment": "john snow"})
	s.Require().NoError(err)

	_, err = s.coll.InsertOne(ctx, bson.M{"_id": 3, "name": "itssnowing", "comment": "ryan gosling"})
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TearDownSuite() {
	ctx := context.Background()

	err := s.db.Drop(ctx)
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TestGet() {
	ctx := context.Background()
	repo := RepoMongo{coll: s.coll}

	want := &User{
		ID:      1,
		Name:    "John Snow",
		Comment: "snow warrior",
	}

	got, err := repo.Get(ctx, 1)
	s.Require().NoError(err)
	s.Equal(want, got)
}

func (s *RepoTestSuite) TestUpsert() {
	ctx := context.Background()
	repo := RepoMongo{coll: s.coll}

	user := User{
		ID:      42,
		Name:    "James Bond",
		Comment: "007",
	}

	err := repo.Upsert(ctx, user)
	s.Require().NoError(err)
}

func (s *RepoTestSuite) TestUpdate() {
	ctx := context.Background()
	repo := RepoMongo{coll: s.coll}

	user := User{
		ID:      1,
		Name:    "Bruce Lee",
		Comment: "cool",
	}

	err := repo.Upsert(ctx, user)
	s.Require().NoError(err)
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}

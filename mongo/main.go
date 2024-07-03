package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const documents = 50000

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	db := newDbConn(ctx)
	defer db.client.Disconnect(ctx)

	//err := db.coll.Drop(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//db.insertManyDocs(ctx)
	////db.updateManyDocs(ctx)
	db.findAndStoreIds(ctx)

}

type productDoc struct {
	ID   uuid.UUID `bson:"_id"`
	Name string    `bson:"name"`
}

type productDocs []productDoc

type localityDoc struct {
	ID     uuid.UUID `bson:"_id"`
	Type   string    `bson:"type"`
	Active bool      `bson:"active"`
}

type localityDocs []localityDoc

type rubricDoc struct {
	ID uuid.UUID `bson:"_id"`
}

type rubricDocs []rubricDoc

type dbConn struct {
	client *mongo.Client
	db     *mongo.Database
	coll   *mongo.Collection
}

func newDbConn(ctx context.Context) dbConn {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetRegistry(BsonRegistry()))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("promo-testing")
	coll := db.Collection("product")

	return dbConn{
		client: client,
		db:     db,
		coll:   coll,
	}

}

func (db *dbConn) findAndStoreIds(ctx context.Context) {
	curr, err := db.coll.Find(ctx, bson.D{{"state", "active"}, {"promoProperties", bson.D{{"$exists", false}}}})
	if err != nil {
		log.Fatal(err)
	}

	var results productDocs

	if err = curr.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	output, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()

	for i, result := range results {
		if i == 5000 {
			break
		}

		id := fmt.Sprintf("%q,", result.ID.String())
		//id := result.ID.String()
		_, err = output.WriteString(id + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (db *dbConn) insertManyDocs(ctx context.Context) {
	docs := prepareDocs()

	start := time.Now()
	_, err := db.coll.InsertMany(ctx, docs)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("INSERT MANY TIME: ", -start.Sub(time.Now()))
}

func (db *dbConn) updateManyDocs(ctx context.Context) {
	docs := prepareDocs()
	prDocs := make(productDocs, 0, len(docs))

	for _, v := range docs {
		c := v.(productDoc)

		prDocs = append(prDocs, c)
	}

	models := mongoWriteModels(prDocs)

	start := time.Now()
	_, err := db.coll.BulkWrite(ctx, models)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("BULK WRITE TIME: ", -start.Sub(time.Now()))
}

func prepareDocs() []interface{} {
	docs := make([]interface{}, 0, documents)

	for i := 0; i < documents; i++ {
		doc := productDoc{
			ID:   uuid.New(),
			Name: fmt.Sprintf("Name_%d", i),
		}

		docs = append(docs, doc)
	}

	return docs
}

func mongoWriteModels(docs productDocs) []mongo.WriteModel {
	models := make([]mongo.WriteModel, 0, len(docs))

	for _, v := range docs {
		filter := bson.M{"_id": v.ID}
		model := mongo.NewUpdateManyModel().
			SetFilter(filter).
			SetUpsert(true).
			SetUpdate(bson.M{"$set": bson.M{"_id": v.ID, "name": v.Name}})

		models = append(models, model)
	}

	return models
}

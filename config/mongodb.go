package config

import (
	"EnglishVocab/pkg/errors"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"os"
	"time"
)

func OpenNewMongoDBConnection(ctx context.Context) (client *mongo.Client, db *mongo.Database, err error) {
	client, err = mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_HOST")))
	if err != nil {
		return nil, nil, errors.New("Error connecting to MongoDB: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db = client.Database(os.Getenv("MONGO_DATABASE"))
	return client, db, nil
}

package main

import (
	"EnglishVocab/config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	ctx := context.Background()
	r := gin.Default()

	// open connection to mongodb
	mongoClient, mongoDB, err := config.OpenNewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	defer func() {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			log.Fatal("Error disconnecting from MongoDB:", err)
		}
	}()

	Router(r, mongoDB)

	log.Println("Starting server on port", os.Getenv("PORT"))

	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("server start error", err)
	}
}

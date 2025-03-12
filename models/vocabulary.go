package models

import (
	"EnglishVocab/models/dto"
	"EnglishVocab/pkg/errors"
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"time"
)

type VocabularyRepository interface {
	SaveNewVocabulary(ctx context.Context, request *dto.VocabularyRequest) error
}

type VocabularyDBRepository struct {
	mongoDB *mongo.Database
}

func NewVocabularyRepository(mongoDB *mongo.Database) VocabularyRepository {
	return &VocabularyDBRepository{mongoDB: mongoDB}
}

func (v VocabularyDBRepository) SaveNewVocabulary(ctx context.Context, request *dto.VocabularyRequest) error {
	if v.mongoDB == nil {
		return errors.New("error on save new vocabulary: ", "MongoDB client is nil")
	}

	request.CreatedAt = time.Now()

	result, err := v.mongoDB.Collection("vocabularies").InsertOne(ctx, request)
	if err != nil {
		return errors.New("error on save new vocabulary: ", err.Error())
	}

	log.Println("Successfully inserted new vocabulary with id:", result.InsertedID)
	return nil
}

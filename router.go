package main

import (
	"EnglishVocab/models"
	"EnglishVocab/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
)

func Router(r *gin.Engine, mongoDB *mongo.Database) {
	vocabularyService := services.NewVocabularyService(models.NewVocabularyRepository(mongoDB))

	api := r.Group("/api")
	v1 := api.Group("/v1")

	healthCheck(v1)

	vocabularyGroup := v1.Group("/vocabulary")
	vocabulary(vocabularyGroup, vocabularyService)

}

func healthCheck(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func vocabulary(r *gin.RouterGroup, service services.VocabularyService) {
	r.POST("", service.SaveNewVocabulary)
}

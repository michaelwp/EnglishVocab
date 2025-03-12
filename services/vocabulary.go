package services

import (
	"EnglishVocab/models"
	"EnglishVocab/models/dto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type VocabularyService interface {
	SaveNewVocabulary(g *gin.Context)
}

type VocabularyRepositoryService struct {
	repo models.VocabularyRepository
}

func NewVocabularyService(repo models.VocabularyRepository) VocabularyService {
	return &VocabularyRepositoryService{repo: repo}
}

func (v VocabularyRepositoryService) SaveNewVocabulary(g *gin.Context) {
	var request *dto.VocabularyRequest

	err := g.ShouldBindJSON(&request)
	if err != nil {
		log.Println("error on binding json vocabulary request:", err)
		g.JSON(http.StatusBadRequest, JsonResponse{
			Status:      ERROR,
			Message:     "error on binding json vocabulary request",
			Translation: "failed.to.save.new.vocabulary",
		})

		return
	}

	err = v.repo.SaveNewVocabulary(g, request)
	if err != nil {
		log.Println("error on saving new vocabulary:", err)
		g.JSON(http.StatusInternalServerError, JsonResponse{
			Status:      ERROR,
			Message:     "error on saving new vocabulary",
			Translation: "failed.to.save.new.vocabulary",
		})

		return
	}

	g.JSON(http.StatusCreated, JsonResponse{
		Status:      SUCCESS,
		Message:     "successfully saved new vocabulary",
		Translation: "successfully.saved.new.vocabulary",
	})
}

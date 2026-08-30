package application

import (
	"markitos-it-svc-articles/internal/domain/models"
	"markitos-it-svc-articles/internal/domain/persistence"
	"markitos-it-svc-articles/internal/domain/types"
)

type GetArticleUseCase struct {
	repo persistence.ArticleRepository
}

func NewGetArticleUseCase(repo persistence.ArticleRepository) *GetArticleUseCase {
	return &GetArticleUseCase{repo: repo}
}

func (uc *GetArticleUseCase) Get(id string) (*models.Article, error) {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return nil, err
	}

	article, err := uc.repo.Get(validID)
	if err != nil {
		return nil, err
	}

	return article, nil
}

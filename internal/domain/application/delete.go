package application

import (
	"markitos-it-svc-articles/internal/domain/persistence"
	"markitos-it-svc-articles/internal/domain/types"
)

type DeleteArticleUseCase struct {
	repo persistence.ArticleRepository
}

func NewDeleteArticleUseCase(repo persistence.ArticleRepository) *DeleteArticleUseCase {
	return &DeleteArticleUseCase{repo: repo}
}

func (uc *DeleteArticleUseCase) Delete(id string) error {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return err
	}

	err = uc.repo.Delete(validID)
	if err != nil {
		return err
	}

	return nil
}

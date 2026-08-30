package application

import (
	"markitos-it-svc-articles/internal/domain/models"
	"markitos-it-svc-articles/internal/domain/persistence"
)

type SaveArticleUseCase struct {
	repo persistence.ArticleRepository
}

func NewSaveArticleUseCase(repo persistence.ArticleRepository) *SaveArticleUseCase {

	return &SaveArticleUseCase{repo: repo}

}

func (uc *SaveArticleUseCase) Save(title, content string, tags []string) (string, error) {
	article, err := models.NewArticle(title, content, tags)
	if err != nil {
		return "", err
	}

	err = uc.repo.Save(article)
	if err != nil {
		return "", err
	}

	return article.ID.Value(), nil
}

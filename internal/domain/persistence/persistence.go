package persistence

import (
	"markitos-it-svc-articles/internal/domain/models"
	"markitos-it-svc-articles/internal/domain/types"
)

type ArticleRepository interface {
	Save(article *models.Article) error
	Get(id *types.ID) (*models.Article, error)
	Delete(id *types.ID) error
	Update(article *models.Article) error
	List() ([]*models.Article, error)
}

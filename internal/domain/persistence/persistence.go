package persistence

import (
	"markitos-it-svc-faqs/internal/domain/models"
	"markitos-it-svc-faqs/internal/domain/types"
)

type FaqRepository interface {
	Save(faq *models.Faq) error
	Get(id *types.ID) (*models.Faq, error)
	Delete(id *types.ID) error
	Update(faq *models.Faq) error
	List() ([]*models.Faq, error)
}

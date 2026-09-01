package application

import (
	"markitos-it-svc-faqs/internal/domain/models"
	"markitos-it-svc-faqs/internal/domain/persistence"
)

type ListFaqUseCase struct {
	repository persistence.FaqRepository
}

func NewListFaqUseCase(repo persistence.FaqRepository) *ListFaqUseCase {
	return &ListFaqUseCase{repository: repo}
}

func (uc *ListFaqUseCase) List() ([]*models.Faq, error) {
	faqs, err := uc.repository.List()
	if err != nil {
		return nil, err
	}

	return faqs, nil
}

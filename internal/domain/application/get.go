package application

import (
	"markitos-it-svc-faqs/internal/domain/models"
	"markitos-it-svc-faqs/internal/domain/persistence"
	"markitos-it-svc-faqs/internal/domain/types"
)

type GetFaqUseCase struct {
	repo persistence.FaqRepository
}

func NewGetFaqUseCase(repo persistence.FaqRepository) *GetFaqUseCase {
	return &GetFaqUseCase{repo: repo}
}

func (uc *GetFaqUseCase) Get(id string) (*models.Faq, error) {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return nil, err
	}

	faq, err := uc.repo.Get(validID)
	if err != nil {
		return nil, err
	}

	return faq, nil
}

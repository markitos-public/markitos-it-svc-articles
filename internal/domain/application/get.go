package application

import (
	"markitos-it-svc-faqs/internal/domain/models"
	"markitos-it-svc-faqs/internal/domain/persistence"
	"markitos-it-svc-faqs/internal/domain/types"
)

type GetFaqUseCase struct {
	repository persistence.FaqRepository
}

func NewGetFaqUseCase(repo persistence.FaqRepository) *GetFaqUseCase {
	return &GetFaqUseCase{repository: repo}
}

func (uc *GetFaqUseCase) Get(id string) (*models.Faq, error) {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return nil, err
	}

	faq, err := uc.repository.Get(validID)
	if err != nil {
		return nil, err
	}

	return faq, nil
}

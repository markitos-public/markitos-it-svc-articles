package application

import (
	"markitos-it-svc-faqs/internal/domain/persistence"
	"markitos-it-svc-faqs/internal/domain/types"
)

type DeleteFaqUseCase struct {
	repository persistence.FaqRepository
}

func NewDeleteFaqUseCase(repo persistence.FaqRepository) *DeleteFaqUseCase {
	return &DeleteFaqUseCase{repository: repo}
}

func (uc *DeleteFaqUseCase) Delete(id string) error {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return err
	}

	err = uc.repository.Delete(validID)
	if err != nil {
		return err
	}

	return nil
}

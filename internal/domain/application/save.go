package application

import (
	"markitos-it-svc-faqs/internal/domain/models"
	"markitos-it-svc-faqs/internal/domain/persistence"
)

type SaveFaqUseCase struct {
	repository persistence.FaqRepository
}

func NewSaveFaqUseCase(repo persistence.FaqRepository) *SaveFaqUseCase {

	return &SaveFaqUseCase{repository: repo}

}

func (uc *SaveFaqUseCase) Save(title, content string, tags []string) (string, error) {
	faq, err := models.NewFaq(title, content, tags)
	if err != nil {
		return "", err
	}

	err = uc.repository.Save(faq)
	if err != nil {
		return "", err
	}

	return faq.ID.Value(), nil
}

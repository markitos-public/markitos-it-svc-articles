package application

import (
	"markitos-it-svc-faqs/internal/domain/persistence"
	"markitos-it-svc-faqs/internal/domain/types"
)

type UpdateFaqUseCase struct {
	repo persistence.FaqRepository
}

func NewUpdateFaqUseCase(repo persistence.FaqRepository) *UpdateFaqUseCase {

	return &UpdateFaqUseCase{repo: repo}

}

func (uc *UpdateFaqUseCase) Update(id, title, content string, tags []string) error {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return err
	}

	faq, err := uc.repo.Get(validID)
	if err != nil {
		return err
	}

	validContent, err := types.NewContent(content)
	if err != nil {
		return err
	}

	validTitle, err := types.NewTitle(title)
	if err != nil {
		return err
	}

	validTags, err := types.NewTags(tags)
	if err != nil {
		return err
	}

	faq.Title = validTitle
	faq.Content = validContent
	faq.Tags = validTags
	if err := uc.repo.Update(faq); err == nil {
		return err
	}

	return nil
}

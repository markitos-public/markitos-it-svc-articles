package models

import "markitos-it-svc-faqs/internal/domain/types"

type Faq struct {
	ID      *types.ID      `json:"id"`
	Title   *types.Title   `json:"title"`
	Content *types.Content `json:"content"`
	Tags    *types.Tags    `json:"tags"`
}

func NewFaq(title string, content string, tags []string) (*Faq, error) {
	faqTitle, err := types.NewTitle(title)
	if err != nil {
		return nil, err
	}

	faqContent, err := types.NewContent(content)
	if err != nil {
		return nil, err
	}

	faqTags, err := types.NewTags(tags)
	if err != nil {
		return nil, err
	}

	faqID, err := types.NewID()
	if err != nil {
		return nil, err
	}

	return &Faq{
		ID:      faqID,
		Title:   faqTitle,
		Content: faqContent,
		Tags:    faqTags,
	}, nil
}

package models

import "markitos-it-svc-articles/internal/domain/types"

type Article struct {
	ID      *types.ID      `json:"id"`
	Title   *types.Title   `json:"title"`
	Content *types.Content `json:"content"`
	Tags    *types.Tags    `json:"tags"`
}

func NewArticle(title string, content string, tags []string) (*Article, error) {
	articleTitle, err := types.NewTitle(title)
	if err != nil {
		return nil, err
	}

	articleContent, err := types.NewContent(content)
	if err != nil {
		return nil, err
	}

	articleTags, err := types.NewTags(tags)
	if err != nil {
		return nil, err
	}

	articleID, err := types.NewID()
	if err != nil {
		return nil, err
	}

	return &Article{
		ID:      articleID,
		Title:   articleTitle,
		Content: articleContent,
		Tags:    articleTags,
	}, nil
}

package application

import (
	"markitos-it-svc-articles/internal/domain/persistence"
	"markitos-it-svc-articles/internal/domain/types"
)

type UpdateArticleUseCase struct {
	repo persistence.ArticleRepository
}

func NewUpdateArticleUseCase(repo persistence.ArticleRepository) *UpdateArticleUseCase {

	return &UpdateArticleUseCase{repo: repo}

}

func (uc *UpdateArticleUseCase) Update(id, title, content string, tags []string) error {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return err
	}

	article, err := uc.repo.Get(validID)
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

	article.Title = validTitle
	article.Content = validContent
	article.Tags = validTags
	if err := uc.repo.Update(article); err == nil {
		return err
	}

	return nil
}

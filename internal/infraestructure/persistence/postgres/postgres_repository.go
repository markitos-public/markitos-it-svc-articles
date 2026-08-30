package postgres

import (
	"errors"
	"markitos-it-svc-articles/internal/domain/models"
	"markitos-it-svc-articles/internal/domain/persistence"
	"markitos-it-svc-articles/internal/domain/types"

	"gorm.io/gorm"
)

type PostgresArticleRepository struct {
	db *gorm.DB
}

type articleRecord struct {
	ID      string   `gorm:"primaryKey;type:uuid;not null"`
	Title   string   `gorm:"not null"`
	Content string   `gorm:"not null"`
	Tags    []string `gorm:"serializer:json;not null"`
}

func (articleRecord) TableName() string {
	return "articles"
}

func NewPostgresArticleRepository(db *gorm.DB) persistence.ArticleRepository {
	if db == nil {
		panic("postgres repository requires a valid gorm db instance")
	}

	if err := db.AutoMigrate(&articleRecord{}); err != nil {
		panic(err)
	}

	return &PostgresArticleRepository{db: db}
}

func (r *PostgresArticleRepository) Save(article *models.Article) error {
	if article == nil {
		return errors.New("article is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	if article.ID == nil || article.Title == nil || article.Content == nil || article.Tags == nil {
		return errors.New("article is incomplete")
	}

	payload := articleRecord{
		ID:      article.ID.Value(),
		Title:   article.Title.Value(),
		Content: article.Content.Value(),
		Tags:    article.Tags.Value(),
	}

	return r.db.Save(&payload).Error
}

func (r *PostgresArticleRepository) Get(id *types.ID) (*models.Article, error) {
	if id == nil {
		return nil, errors.New("article id is nil")
	}
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	var payload articleRecord
	result := r.db.Where("id = ?", id.Value()).First(&payload)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}

	articleID, err := types.NewIDFromString(payload.ID)
	if err != nil {
		return nil, err
	}

	articleTitle, err := types.NewTitle(payload.Title)
	if err != nil {
		return nil, err
	}

	articleContent, err := types.NewContent(payload.Content)
	if err != nil {
		return nil, err
	}

	articleTags, err := types.NewTags(payload.Tags)
	if err != nil {
		return nil, err
	}

	return &models.Article{
		ID:      articleID,
		Title:   articleTitle,
		Content: articleContent,
		Tags:    articleTags,
	}, nil
}

func (r *PostgresArticleRepository) Delete(id *types.ID) error {
	if id == nil {
		return errors.New("article id is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}

	result := r.db.Delete(&articleRecord{}, "id = ?", id.Value())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PostgresArticleRepository) Update(article *models.Article) error {
	if article == nil {
		return errors.New("article is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	if article.ID == nil || article.Title == nil || article.Content == nil || article.Tags == nil {
		return errors.New("article is incomplete")
	}

	payload := articleRecord{
		ID:      article.ID.Value(),
		Title:   article.Title.Value(),
		Content: article.Content.Value(),
		Tags:    article.Tags.Value(),
	}

	result := r.db.Model(&articleRecord{}).Where("id = ?", payload.ID).Updates(payload)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PostgresArticleRepository) List() ([]*models.Article, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	var payloads []articleRecord
	if err := r.db.Find(&payloads).Error; err != nil {
		return nil, err
	}

	articles := make([]*models.Article, 0, len(payloads))
	for _, payload := range payloads {
		articleID, err := types.NewIDFromString(payload.ID)
		if err != nil {
			return nil, err
		}

		articleTitle, err := types.NewTitle(payload.Title)
		if err != nil {
			return nil, err
		}

		articleContent, err := types.NewContent(payload.Content)
		if err != nil {
			return nil, err
		}

		articleTags, err := types.NewTags(payload.Tags)
		if err != nil {
			return nil, err
		}

		articles = append(articles, &models.Article{
			ID:      articleID,
			Title:   articleTitle,
			Content: articleContent,
			Tags:    articleTags,
		})
	}

	return articles, nil
}

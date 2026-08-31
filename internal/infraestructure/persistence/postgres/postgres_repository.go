package postgres

import (
	"errors"
	"markitos-it-svc-faqs/internal/domain/models"
	"markitos-it-svc-faqs/internal/domain/persistence"
	"markitos-it-svc-faqs/internal/domain/types"

	"gorm.io/gorm"
)

type PostgresFaqRepository struct {
	db *gorm.DB
}

type faqRecord struct {
	ID      string   `gorm:"primaryKey;type:uuid;not null"`
	Title   string   `gorm:"not null"`
	Content string   `gorm:"not null"`
	Tags    []string `gorm:"serializer:json;not null"`
}

func (faqRecord) TableName() string {
	return "faqs"
}

func NewPostgresFaqRepository(db *gorm.DB) persistence.FaqRepository {
	if db == nil {
		panic("postgres repository requires a valid gorm db instance")
	}

	if err := db.AutoMigrate(&faqRecord{}); err != nil {
		panic(err)
	}

	return &PostgresFaqRepository{db: db}
}

func (r *PostgresFaqRepository) Save(faq *models.Faq) error {
	if faq == nil {
		return errors.New("faq is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	if faq.ID == nil || faq.Title == nil || faq.Content == nil || faq.Tags == nil {
		return errors.New("faq is incomplete")
	}

	payload := faqRecord{
		ID:      faq.ID.Value(),
		Title:   faq.Title.Value(),
		Content: faq.Content.Value(),
		Tags:    faq.Tags.Value(),
	}

	return r.db.Save(&payload).Error
}

func (r *PostgresFaqRepository) Get(id *types.ID) (*models.Faq, error) {
	if id == nil {
		return nil, errors.New("faq id is nil")
	}
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	var payload faqRecord
	result := r.db.Where("id = ?", id.Value()).First(&payload)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}

	faqID, err := types.NewIDFromString(payload.ID)
	if err != nil {
		return nil, err
	}

	faqTitle, err := types.NewTitle(payload.Title)
	if err != nil {
		return nil, err
	}

	faqContent, err := types.NewContent(payload.Content)
	if err != nil {
		return nil, err
	}

	faqTags, err := types.NewTags(payload.Tags)
	if err != nil {
		return nil, err
	}

	return &models.Faq{
		ID:      faqID,
		Title:   faqTitle,
		Content: faqContent,
		Tags:    faqTags,
	}, nil
}

func (r *PostgresFaqRepository) Delete(id *types.ID) error {
	if id == nil {
		return errors.New("faq id is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}

	result := r.db.Delete(&faqRecord{}, "id = ?", id.Value())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PostgresFaqRepository) Update(faq *models.Faq) error {
	if faq == nil {
		return errors.New("faq is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	if faq.ID == nil || faq.Title == nil || faq.Content == nil || faq.Tags == nil {
		return errors.New("faq is incomplete")
	}

	payload := faqRecord{
		ID:      faq.ID.Value(),
		Title:   faq.Title.Value(),
		Content: faq.Content.Value(),
		Tags:    faq.Tags.Value(),
	}

	result := r.db.Model(&faqRecord{}).Where("id = ?", payload.ID).Updates(payload)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PostgresFaqRepository) List() ([]*models.Faq, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	var payloads []faqRecord
	if err := r.db.Find(&payloads).Error; err != nil {
		return nil, err
	}

	faqs := make([]*models.Faq, 0, len(payloads))
	for _, payload := range payloads {
		faqID, err := types.NewIDFromString(payload.ID)
		if err != nil {
			return nil, err
		}

		faqTitle, err := types.NewTitle(payload.Title)
		if err != nil {
			return nil, err
		}

		faqContent, err := types.NewContent(payload.Content)
		if err != nil {
			return nil, err
		}

		faqTags, err := types.NewTags(payload.Tags)
		if err != nil {
			return nil, err
		}

		faqs = append(faqs, &models.Faq{
			ID:      faqID,
			Title:   faqTitle,
			Content: faqContent,
			Tags:    faqTags,
		})
	}

	return faqs, nil
}

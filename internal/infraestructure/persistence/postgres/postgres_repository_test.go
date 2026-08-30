package postgres_test

import (
	"testing"

	"markitos-it-svc-articles/internal/domain/models"
	"markitos-it-svc-articles/internal/domain/types"
	postgresrepo "markitos-it-svc-articles/internal/infraestructure/persistence/postgres"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPostgresRepository_SaveAndGet(t *testing.T) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	sqlDB, err := db.DB()
	require.NoError(t, err)

	repo := postgresrepo.NewPostgresArticleRepository(db)

	id, _ := types.NewID()
	title, _ := types.NewTitle("Test Title")
	content, _ := types.NewContent("This is a test content with more than ten characters.")
	tags, _ := types.NewTags([]string{"test", "go"})

	article := &models.Article{
		ID:      id,
		Title:   title,
		Content: content,
		Tags:    tags,
	}

	t.Cleanup(func() {
		err := db.Exec("DELETE FROM articles WHERE id = ?", article.ID.Value()).Error
		require.NoError(t, err)
		require.NoError(t, sqlDB.Close())
	})

	err = repo.Save(article)
	require.NoError(t, err)

	retrievedArticle, err := repo.Get(id)
	require.NoError(t, err)
	assert.Equal(t, article.ID.Value(), retrievedArticle.ID.Value())
	assert.Equal(t, article.Title.Value(), retrievedArticle.Title.Value())
	assert.Equal(t, article.Content.Value(), retrievedArticle.Content.Value())
	assert.Equal(t, article.Tags.Value(), retrievedArticle.Tags.Value())
}

func TestPostgresRepository_Save_NilArticle(t *testing.T) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	repo := postgresrepo.NewPostgresArticleRepository(db)

	err = repo.Save(nil)
	assert.Error(t, err)
	assert.Equal(t, "article is nil", err.Error())
}

func TestPostgresRepository_Save_IncompleteArticle(t *testing.T) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	repo := postgresrepo.NewPostgresArticleRepository(db)

	id, _ := types.NewID()
	title, _ := types.NewTitle("Test Title")
	content, _ := types.NewContent("This is a test content with more than ten characters.")

	article := &models.Article{
		ID:      id,
		Title:   title,
		Content: content,
		Tags:    nil, // Tags is nil
	}

	err = repo.Save(article)
	assert.Error(t, err)
	assert.Equal(t, "article is incomplete", err.Error())
}

func TestPostgresRepository_Get_NonExistentArticle(t *testing.T) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	repo := postgresrepo.NewPostgresArticleRepository(db)

	nonExistentID, _ := types.NewID()

	article, err := repo.Get(nonExistentID)
	assert.Error(t, err)
	assert.Nil(t, article)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestPostgresRepository_DeleteExistingArticle(t *testing.T) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	repo := postgresrepo.NewPostgresArticleRepository(db)

	id, _ := types.NewID()
	title, _ := types.NewTitle("Test Title")
	content, _ := types.NewContent("This is a test content with more than ten characters.")
	tags, _ := types.NewTags([]string{"test", "go"})

	article := &models.Article{
		ID:      id,
		Title:   title,
		Content: content,
		Tags:    tags,
	}

	err = repo.Save(article)
	require.NoError(t, err)

	err = repo.Delete(id)
	require.NoError(t, err)

	deletedArticle, err := repo.Get(id)
	assert.Error(t, err)
	assert.Nil(t, deletedArticle)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestPostgresRepository_DeleteNonExistentArticle(t *testing.T) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	repo := postgresrepo.NewPostgresArticleRepository(db)

	nonExistentID, _ := types.NewID()

	err = repo.Delete(nonExistentID)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestPostgresRepository_Delete_NilID(t *testing.T) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	repo := postgresrepo.NewPostgresArticleRepository(db)

	err = repo.Delete(nil)
	assert.Error(t, err)
	assert.Equal(t, "article id is nil", err.Error())
}

func TestPostgresRepository_UpdateExistingArticle(t *testing.T) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	repo := postgresrepo.NewPostgresArticleRepository(db)

	id, _ := types.NewID()
	title, _ := types.NewTitle("Test Title")
	content, _ := types.NewContent("This is a test content with more than ten characters.")
	tags, _ := types.NewTags([]string{"test", "go"})

	article := &models.Article{
		ID:      id,
		Title:   title,
		Content: content,
		Tags:    tags,
	}

	err = repo.Save(article)
	require.NoError(t, err)

	newTitle, _ := types.NewTitle("Updated Title")
	newContent, _ := types.NewContent("This is updated content with more than ten characters.")
	newTags, _ := types.NewTags([]string{"updated", "go"})

	article.Title = newTitle
	article.Content = newContent
	article.Tags = newTags

	err = repo.Update(article)
	require.NoError(t, err)

	updatedArticle, err := repo.Get(id)
	require.NoError(t, err)
	assert.Equal(t, article.ID.Value(), updatedArticle.ID.Value())
	assert.Equal(t, article.Title.Value(), updatedArticle.Title.Value())
	assert.Equal(t, article.Content.Value(), updatedArticle.Content.Value())
	assert.Equal(t, article.Tags.Value(), updatedArticle.Tags.Value())
}

func TestPostgresRepository_List(t *testing.T) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	sqlDB, err := db.DB()
	require.NoError(t, err)

	repo := postgresrepo.NewPostgresArticleRepository(db)

	id1, _ := types.NewID()
	title1, _ := types.NewTitle("First Test Title")
	content1, _ := types.NewContent("This is the first test content with enough characters.")
	tags1, _ := types.NewTags([]string{"first", "test"})

	article1 := &models.Article{
		ID:      id1,
		Title:   title1,
		Content: content1,
		Tags:    tags1,
	}

	id2, _ := types.NewID()
	title2, _ := types.NewTitle("Second Test Title")
	content2, _ := types.NewContent("This is the second test content with enough characters.")
	tags2, _ := types.NewTags([]string{"second", "test"})

	article2 := &models.Article{
		ID:      id2,
		Title:   title2,
		Content: content2,
		Tags:    tags2,
	}

	t.Cleanup(func() {
		err := db.Exec("DELETE FROM articles WHERE id IN (?, ?)", article1.ID.Value(), article2.ID.Value()).Error
		require.NoError(t, err)
		require.NoError(t, sqlDB.Close())
	})

	err = repo.Save(article1)
	require.NoError(t, err)

	err = repo.Save(article2)
	require.NoError(t, err)

	articles, err := repo.List()
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(articles), 2)

	found1 := false
	found2 := false
	for _, a := range articles {
		if a.ID.Value() == article1.ID.Value() {
			found1 = true
			assert.Equal(t, article1.Title.Value(), a.Title.Value())
			assert.Equal(t, article1.Content.Value(), a.Content.Value())
			assert.Equal(t, article1.Tags.Value(), a.Tags.Value())
		}
		if a.ID.Value() == article2.ID.Value() {
			found2 = true
			assert.Equal(t, article2.Title.Value(), a.Title.Value())
			assert.Equal(t, article2.Content.Value(), a.Content.Value())
			assert.Equal(t, article2.Tags.Value(), a.Tags.Value())
		}
	}

	assert.True(t, found1, "article1 should be present in the list")
	assert.True(t, found2, "article2 should be present in the list")
}

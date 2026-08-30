package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-it-svc-articles/internal/infraestructure/persistence/postgres"
	"markitos-it-svc-articles/internal/infraestructure/rest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type articleDTO struct {
	ID      string   `json:"id,omitempty"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func setupTestRouter(t *testing.T) (*httptest.Server, *gorm.DB) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	sqlDB, err := db.DB()
	require.NoError(t, err)

	repo := postgres.NewPostgresArticleRepository(db)

	saveHandler := rest.NewRESTSaveUseCase(repo)
	getHandler := rest.NewRESTGetUseCase(repo)
	updateHandler := rest.NewRESTUpdateUseCase(repo)
	deleteHandler := rest.NewRESTDeleteUseCase(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /articles", saveHandler.Save)
	mux.HandleFunc("GET /articles/{id}", getHandler.Get)
	mux.HandleFunc("PUT /articles/{id}", updateHandler.Update)
	mux.HandleFunc("DELETE /articles/{id}", deleteHandler.Delete)

	server := httptest.NewServer(mux)

	t.Cleanup(func() {
		db.Exec("DELETE FROM articles")
		sqlDB.Close()
		server.Close()
	})

	return server, db
}

func TestREST_ArticleLifecycle(t *testing.T) {
	server, _ := setupTestRouter(t)

	newArticle := articleDTO{
		Title:   "REST API Title",
		Content: "This is a valid content string for testing REST handlers.",
		Tags:    []string{"rest", "api"},
	}

	body, err := json.Marshal(newArticle)
	require.NoError(t, err)

	resp, err := http.Post(server.URL+"/articles", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createdDTO articleDTO
	err = json.NewDecoder(resp.Body).Decode(&createdDTO)
	require.NoError(t, err)
	assert.NotEmpty(t, createdDTO.ID)
	assert.Equal(t, newArticle.Title, createdDTO.Title)

	getResp, err := http.Get(server.URL + "/articles/" + createdDTO.ID)
	require.NoError(t, err)
	defer getResp.Body.Close()

	assert.Equal(t, http.StatusOK, getResp.StatusCode)
	var fetchedDTO articleDTO
	err = json.NewDecoder(getResp.Body).Decode(&fetchedDTO)
	require.NoError(t, err)
	assert.Equal(t, createdDTO.ID, fetchedDTO.ID)

	updatedArticle := articleDTO{
		Title:   "Updated REST Title",
		Content: "This is updated valid content string for testing REST handlers.",
		Tags:    []string{"rest", "updated"},
	}
	updateBody, err := json.Marshal(updatedArticle)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, server.URL+"/articles/"+createdDTO.ID, bytes.NewBuffer(updateBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	updateResp, err := client.Do(req)
	require.NoError(t, err)
	defer updateResp.Body.Close()

	assert.Equal(t, http.StatusOK, updateResp.StatusCode)

	delReq, err := http.NewRequest(http.MethodDelete, server.URL+"/articles/"+createdDTO.ID, nil)
	require.NoError(t, err)

	delResp, err := client.Do(delReq)
	require.NoError(t, err)
	defer delResp.Body.Close()

	assert.Equal(t, http.StatusNoContent, delResp.StatusCode)
}

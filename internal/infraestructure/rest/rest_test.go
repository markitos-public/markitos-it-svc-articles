package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-it-svc-faqs/internal/infraestructure/persistence/postgres"
	"markitos-it-svc-faqs/internal/infraestructure/rest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type faqDTO struct {
	ID      string   `json:"id,omitempty"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func setupTestRouter(t *testing.T) (*httptest.Server, *gorm.DB) {
	db, err := gorm.Open(pgdriver.Open("host=localhost port=5431 user=postgres password=postgres dbname=faqs_test sslmode=disable"), &gorm.Config{})
	require.NoError(t, err)

	sqlDB, err := db.DB()
	require.NoError(t, err)

	repo := postgres.NewPostgresFaqRepository(db)

	saveHandler := rest.NewRESTSaveUseCase(repo)
	getHandler := rest.NewRESTGetUseCase(repo)
	updateHandler := rest.NewRESTUpdateUseCase(repo)
	deleteHandler := rest.NewRESTDeleteUseCase(repo)
	listHandler := rest.NewRESTListUseCase(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /faqs", saveHandler.Save)
	mux.HandleFunc("GET /faqs/{id}", getHandler.Get)
	mux.HandleFunc("GET /faqs", listHandler.List)
	mux.HandleFunc("PUT /faqs/{id}", updateHandler.Update)
	mux.HandleFunc("DELETE /faqs/{id}", deleteHandler.Delete)

	server := httptest.NewServer(mux)

	t.Cleanup(func() {
		db.Exec("DELETE FROM faqs")
		sqlDB.Close()
		server.Close()
	})

	return server, db
}

func TestREST_FaqLifecycle(t *testing.T) {
	server, _ := setupTestRouter(t)

	newFaq := faqDTO{
		Title:   "REST API Title",
		Content: "This is a valid content string for testing REST handlers.",
		Tags:    []string{"rest", "api"},
	}

	body, err := json.Marshal(newFaq)
	require.NoError(t, err)

	resp, err := http.Post(server.URL+"/faqs", "application/json", bytes.NewBuffer(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createdDTO faqDTO
	err = json.NewDecoder(resp.Body).Decode(&createdDTO)
	require.NoError(t, err)
	assert.NotEmpty(t, createdDTO.ID)
	assert.Equal(t, newFaq.Title, createdDTO.Title)

	getResp, err := http.Get(server.URL + "/faqs/" + createdDTO.ID)
	require.NoError(t, err)
	defer getResp.Body.Close()

	assert.Equal(t, http.StatusOK, getResp.StatusCode)
	var fetchedDTO faqDTO
	err = json.NewDecoder(getResp.Body).Decode(&fetchedDTO)
	require.NoError(t, err)
	assert.Equal(t, createdDTO.ID, fetchedDTO.ID)

	updatedFaq := faqDTO{
		Title:   "Updated REST Title",
		Content: "This is updated valid content string for testing REST handlers.",
		Tags:    []string{"rest", "updated"},
	}
	updateBody, err := json.Marshal(updatedFaq)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPut, server.URL+"/faqs/"+createdDTO.ID, bytes.NewBuffer(updateBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	updateResp, err := client.Do(req)
	require.NoError(t, err)
	defer updateResp.Body.Close()

	assert.Equal(t, http.StatusOK, updateResp.StatusCode)

	delReq, err := http.NewRequest(http.MethodDelete, server.URL+"/faqs/"+createdDTO.ID, nil)
	require.NoError(t, err)

	delResp, err := client.Do(delReq)
	require.NoError(t, err)
	defer delResp.Body.Close()

	assert.Equal(t, http.StatusNoContent, delResp.StatusCode)
}

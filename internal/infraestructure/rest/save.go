package rest

import (
	"encoding/json"
	"markitos-it-svc-articles/internal/domain/application"
	"markitos-it-svc-articles/internal/domain/persistence"
	"net/http"
)

type RESTSaveUseCase struct {
	repository persistence.ArticleRepository
}

func NewRESTSaveUseCase(repository persistence.ArticleRepository) *RESTSaveUseCase {
	return &RESTSaveUseCase{
		repository: repository,
	}
}

func (uc *RESTSaveUseCase) Save(w http.ResponseWriter, r *http.Request) {
	var dto articleDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	saveUC := application.NewSaveArticleUseCase(uc.repository)
	id, err := saveUC.Save(dto.Title, dto.Content, dto.Tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto)
}

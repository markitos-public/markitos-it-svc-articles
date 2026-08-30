package rest

import (
	"encoding/json"
	"markitos-it-svc-articles/internal/domain/persistence"
	"net/http"
)

type RESTListUseCase struct {
	repository persistence.ArticleRepository
}

func NewRESTListUseCase(repository persistence.ArticleRepository) *RESTListUseCase {
	return &RESTListUseCase{
		repository: repository,
	}
}

func (uc *RESTListUseCase) List(w http.ResponseWriter, r *http.Request) {
	articles, err := uc.repository.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]articleDTO, 0, len(articles))
	for _, a := range articles {
		response = append(response, articleDTO{
			ID:      a.ID.Value(),
			Title:   a.Title.Value(),
			Content: a.Content.Value(),
			Tags:    a.Tags.Value(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

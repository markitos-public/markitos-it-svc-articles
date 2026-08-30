package rest

import (
	"encoding/json"
	"markitos-it-svc-articles/internal/domain/application"
	"markitos-it-svc-articles/internal/domain/persistence"
	"net/http"
)

type RESTGetUseCase struct {
	repository persistence.ArticleRepository
}

func NewRESTGetUseCase(repository persistence.ArticleRepository) *RESTGetUseCase {
	return &RESTGetUseCase{
		repository: repository,
	}
}

func (uc *RESTGetUseCase) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	getUC := application.NewGetArticleUseCase(uc.repository)
	article, err := getUC.Get(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := articleDTO{
		ID:      article.ID.Value(),
		Title:   article.Title.Value(),
		Content: article.Content.Value(),
		Tags:    article.Tags.Value(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

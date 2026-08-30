package rest

import (
	"markitos-it-svc-articles/internal/domain/application"
	"markitos-it-svc-articles/internal/domain/persistence"
	"net/http"
)

type RESTDeleteUseCase struct {
	repository persistence.ArticleRepository
}

func NewRESTDeleteUseCase(repository persistence.ArticleRepository) *RESTDeleteUseCase {
	return &RESTDeleteUseCase{
		repository: repository,
	}
}

func (uc *RESTDeleteUseCase) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	deleteUC := application.NewDeleteArticleUseCase(uc.repository)
	if err := deleteUC.Delete(idStr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

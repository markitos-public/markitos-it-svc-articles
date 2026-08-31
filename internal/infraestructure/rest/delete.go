package rest

import (
	"markitos-it-svc-faqs/internal/domain/application"
	"markitos-it-svc-faqs/internal/domain/persistence"
	"net/http"
)

type RESTDeleteUseCase struct {
	repository persistence.FaqRepository
}

func NewRESTDeleteUseCase(repository persistence.FaqRepository) *RESTDeleteUseCase {
	return &RESTDeleteUseCase{
		repository: repository,
	}
}

func (uc *RESTDeleteUseCase) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	deleteUC := application.NewDeleteFaqUseCase(uc.repository)
	if err := deleteUC.Delete(idStr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

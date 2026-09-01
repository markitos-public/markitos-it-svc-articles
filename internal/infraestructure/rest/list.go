package rest

import (
	"encoding/json"
	"markitos-it-svc-faqs/internal/domain/application"
	"markitos-it-svc-faqs/internal/domain/persistence"
	"net/http"
)

type RESTListUseCase struct {
	repository persistence.FaqRepository
}

func NewRESTListUseCase(repository persistence.FaqRepository) *RESTListUseCase {
	return &RESTListUseCase{
		repository: repository,
	}
}

func (uc *RESTListUseCase) List(w http.ResponseWriter, r *http.Request) {
	listUC := application.NewListFaqUseCase(uc.repository)
	faqs, err := listUC.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := make([]faqDTO, 0, len(faqs))
	for _, a := range faqs {
		response = append(response, faqDTO{
			ID:      a.ID.Value(),
			Title:   a.Title.Value(),
			Content: a.Content.Value(),
			Tags:    a.Tags.Value(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

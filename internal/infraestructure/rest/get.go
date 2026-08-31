package rest

import (
	"encoding/json"
	"markitos-it-svc-faqs/internal/domain/application"
	"markitos-it-svc-faqs/internal/domain/persistence"
	"net/http"
)

type RESTGetUseCase struct {
	repository persistence.FaqRepository
}

func NewRESTGetUseCase(repository persistence.FaqRepository) *RESTGetUseCase {
	return &RESTGetUseCase{
		repository: repository,
	}
}

func (uc *RESTGetUseCase) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	getUC := application.NewGetFaqUseCase(uc.repository)
	faq, err := getUC.Get(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := faqDTO{
		ID:      faq.ID.Value(),
		Title:   faq.Title.Value(),
		Content: faq.Content.Value(),
		Tags:    faq.Tags.Value(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

package main

import (
	"log"
	"net/http"

	"markitos-it-svc-articles/internal/infraestructure/persistence/postgres"
	"markitos-it-svc-articles/internal/infraestructure/rest"

	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type articleDTO struct {
	ID      string   `json:"id,omitempty"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func main() {
	db, err := gorm.Open(pgdriver.Open(resolveDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	repo := postgres.NewPostgresArticleRepository(db)
	saveHandler := rest.NewRESTSaveUseCase(repo)
	getHandler := rest.NewRESTGetUseCase(repo)
	updateHandler := rest.NewRESTUpdateUseCase(repo)
	deleteHandler := rest.NewRESTDeleteUseCase(repo)
	listHandler := rest.NewRESTListUseCase(repo)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /articles", func(w http.ResponseWriter, r *http.Request) {
		listHandler.List(w, r)
	})

	mux.HandleFunc("GET /articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		getHandler.Get(w, r)
	})

	mux.HandleFunc("POST /articles", func(w http.ResponseWriter, r *http.Request) {
		saveHandler.Save(w, r)

	})

	mux.HandleFunc("PUT /articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		updateHandler.Update(w, r)
	})

	mux.HandleFunc("DELETE /articles/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteHandler.Delete(w, r)
	})

	address := resolveAddress()
	log.Printf("Starting REST service on port %s...", address)
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func resolveDSN() string {
	return "host=localhost port=5431 user=postgres password=postgres dbname=articles_test sslmode=disable"
}

func resolveAddress() string {
	return ":8080"
}

package main

import (
	"log"
	"net/http"
	"os"

	"markitos-it-svc-faqs/internal/infraestructure/persistence/postgres"
	"markitos-it-svc-faqs/internal/infraestructure/rest"

	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type faqDTO struct {
	ID      string   `json:"id,omitempty"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func main() {
	db, err := connectDatabase()
	if err != nil {
		log.Fatalf("database connectivity check failed: %v", err)
	}

	repo := postgres.NewPostgresFaqRepository(db)
	saveHandler := rest.NewRESTSaveUseCase(repo)
	getHandler := rest.NewRESTGetUseCase(repo)
	updateHandler := rest.NewRESTUpdateUseCase(repo)
	deleteHandler := rest.NewRESTDeleteUseCase(repo)
	listHandler := rest.NewRESTListUseCase(repo)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /faqs", func(w http.ResponseWriter, r *http.Request) {
		listHandler.List(w, r)
	})

	mux.HandleFunc("GET /faqs/{id}", func(w http.ResponseWriter, r *http.Request) {
		getHandler.Get(w, r)
	})

	mux.HandleFunc("POST /faqs", func(w http.ResponseWriter, r *http.Request) {
		saveHandler.Save(w, r)

	})

	mux.HandleFunc("PUT /faqs/{id}", func(w http.ResponseWriter, r *http.Request) {
		updateHandler.Update(w, r)
	})

	mux.HandleFunc("DELETE /faqs/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteHandler.Delete(w, r)
	})

	address := resolveAddress()
	log.Printf("Database connectivity OK. Starting REST service on port %s...", address)
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func connectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(pgdriver.Open(resolveDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, err
	}

	return db, nil
}

func resolveDSN() string {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN must be set")
	}

	return dsn
}

func resolveAddress() string {
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		log.Fatal("SERVER_ADDRESS must be set")
	}

	return address
}

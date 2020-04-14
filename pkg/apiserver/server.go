package apiserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/replicatedhq/troubleshoot-preview/pkg/handlers"
)

func Start() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/healthz", handlers.Healthz)

	r.HandleFunc("/v1/preflight/{id}", handlers.ServePreflight).Methods("GET")
	r.HandleFunc("/v1/preflight/{id}", handlers.UpdatePreflight).Methods("PUT", "OPTIONS")
	r.HandleFunc("/v1/preflight", handlers.CreatePreflight).Methods("POST", "OPTIONS")

	r.HandleFunc("/v1/support-bundle/{id}", handlers.ServeSupportBundle).Methods("GET")
	r.HandleFunc("/v1/support-bundle/{id}", handlers.UpdateSupportBundle).Methods("PUT", "OPTIONS")
	r.HandleFunc("/v1/support-bundle", handlers.CreateSupportBundle).Methods("POST", "OPTIONS")

	srv := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}

	fmt.Printf("Starting troubleshoot-preview API on port %d...\n", 3000)

	log.Fatal(srv.ListenAndServe())
}

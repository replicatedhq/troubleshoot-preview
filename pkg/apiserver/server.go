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
	srv := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}

	fmt.Printf("Starting troubleshoot-preview API on port %d...\n", 3000)

	log.Fatal(srv.ListenAndServe())
}

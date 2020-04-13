package handlers

import (
	"net/http"
)

type HealthzResponse struct {
	Version string `json:"version"`
	GitSHA  string `json:"gitSha"`
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	healthzResponse := HealthzResponse{
		Version: "test",
		GitSHA:  "test",
	}

	statusCode := 200

	JSON(w, statusCode, healthzResponse)
}

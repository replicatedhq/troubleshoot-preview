package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/replicatedhq/troubleshoot-preview/pkg/logger"
	"gopkg.in/yaml.v2"
)

type Handlers struct {
}

func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func YAML(w http.ResponseWriter, code int, payload interface{}) {
	response, err := yaml.Marshal(payload)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/x-yaml")
	w.WriteHeader(code)
	w.Write(response)
}

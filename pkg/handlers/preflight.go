package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/replicatedhq/troubleshoot-preview/pkg/logger"
	"github.com/replicatedhq/troubleshoot-preview/pkg/persistence"
	"github.com/teris-io/shortid"
	"go.uber.org/zap"
)

type CreatePreflightResponse struct {
	ID string `json:"id"`
}

type PreflightRequest struct {
	Spec string `json:"spec"`
}

func CreatePreflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type, origin, accept, authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	logger.Debug("creating preflight")

	preflightRequest := PreflightRequest{}
	if err := json.NewDecoder(r.Body).Decode(&preflightRequest); err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	generator := shortid.MustNew(1, shortid.DefaultABC, uint64(time.Now().UnixNano()))
	id := generator.MustGenerate()

	cmd := persistence.MustGetRedisClient().Set(id, preflightRequest.Spec, 0)
	if cmd.Err() != nil {
		logger.Error(cmd.Err())
		w.WriteHeader(500)
		return
	}

	createPreflightResponse := CreatePreflightResponse{
		ID: id,
	}

	JSON(w, 201, createPreflightResponse)
}

func UpdatePreflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type, origin, accept, authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	logger.Debug("updating preflight",
		zap.String("id", mux.Vars(r)["id"]))

	preflightRequest := PreflightRequest{}
	if err := json.NewDecoder(r.Body).Decode(&preflightRequest); err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	vars := mux.Vars(r)
	cmd := persistence.MustGetRedisClient().Set(vars["id"], preflightRequest.Spec, 0)
	if cmd.Err() != nil {
		logger.Error(cmd.Err())
		w.WriteHeader(500)
		return
	}

	JSON(w, 204, nil)
}

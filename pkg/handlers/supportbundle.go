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

type CreateSupportBundleResponse struct {
	ID string `json:"id"`
}

type SupportBundleRequest struct {
	Spec string `json:"spec"`
}

func CreateSupportBundle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type, origin, accept, authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	logger.Debug("creating support-bundle")

	supportBundleRequest := SupportBundleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&supportBundleRequest); err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	generator := shortid.MustNew(1, shortid.DefaultABC, uint64(time.Now().UnixNano()))
	id := generator.MustGenerate()

	cmd := persistence.MustGetRedisClient().Set(id, supportBundleRequest.Spec, 0)
	if cmd.Err() != nil {
		logger.Error(cmd.Err())
		w.WriteHeader(500)
		return
	}

	createSupportBundleResponse := CreateSupportBundleResponse{
		ID: id,
	}

	JSON(w, 201, createSupportBundleResponse)
}

func UpdateSupportBundle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type, origin, accept, authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	logger.Debug("updating support-bundle",
		zap.String("id", mux.Vars(r)["id"]))

	supportBundleRequest := SupportBundleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&supportBundleRequest); err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		return
	}

	vars := mux.Vars(r)
	cmd := persistence.MustGetRedisClient().Set(vars["id"], supportBundleRequest.Spec, 0)
	if cmd.Err() != nil {
		logger.Error(cmd.Err())
		w.WriteHeader(500)
		return
	}

	JSON(w, 204, nil)
}

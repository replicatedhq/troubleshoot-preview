package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/replicatedhq/troubleshoot-preview/pkg/logger"
	"github.com/replicatedhq/troubleshoot-preview/pkg/persistence"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cmd := persistence.MustGetRedisClient().Get(vars["id"])
	if cmd.Err() != nil {
		logger.Error(cmd.Err())
		w.WriteHeader(500)
		return
	}

	YAML(w, 200, cmd.Val())
}

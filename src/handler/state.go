package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"

	"terraform_http_backend/src"
	"terraform_http_backend/src/log"
	"terraform_http_backend/src/model"
	"terraform_http_backend/src/store"
	"terraform_http_backend/src/store/file"
)

func GetState(w http.ResponseWriter, r *http.Request) {
	projectName := chi.URLParam(r, "projectName")

	store := getStateStore()

	state := store.GetState(projectName)
	if state == nil {
		log.Warn("No state found in the store")
		state = &model.TFState{Version: 1}
	}

	json.NewEncoder(w).Encode(state)
	log.Info("State for project: %s returned", projectName)
}

func SetState(w http.ResponseWriter, r *http.Request) {
	projectName := chi.URLParam(r, "projectName")

	store := getStateStore()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err, "Failed to read body bytes")
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var state model.TFState

	err = json.Unmarshal(body, &state)
	if err != nil {
		log.Error(err, "Failed to unmarshal body into struct")
		http.Error(w, "Invalid request Payload", http.StatusBadRequest)
		return
	}

	err = store.SaveState(projectName, state)
	if err != nil {
		http.Error(w, "Failed to save the state", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Info("State for project: %s saved", projectName)

}

func getStateStore() store.Store {
	conf := src.GetConfig()

	switch conf.StoreType {
	default:
		return file.GetInstance()
	}
}

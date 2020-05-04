package handlers

import (
	"api/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/square/go-jose.v2/json"
)

// HandleInfo returns information about uploaded file
func HandleInfo(env *Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	data := models.Audio{}
	err := env.DB.Read(models.DBAudio, vars["id"], &data)
	if err != nil {
		return HTTPError(http.StatusNotFound, "file_not_found", err)
	}
	data.Views = data.Views + 1
	data.ViewedAt = time.Now().UTC().Format(time.RFC3339)
	env.DB.Write(models.DBAudio, vars["id"], data)

	json, err := json.Marshal(data)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "cant_generate_response", err)
	}

	w.Write(json)

	return nil
}

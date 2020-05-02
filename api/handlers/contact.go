package handlers

import (
	"api/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lithammer/shortuuid/v3"
)

// HandleContact handles requests to /contact endpoint
func HandleContact(env *Env, w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return HTTPError(http.StatusBadRequest, "cant_read_request_body", err)
	}

	model := models.Contact{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return HTTPError(http.StatusBadRequest, "cant_parse_request_body", err)
	}
	model.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	err = env.DB.Write(models.DBContact, shortuuid.New(), model)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "cant_write_data", err)
	}

	return nil
}

package handlers

import (
	"api/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lithammer/shortuuid/v3"
)

// HandleAbuse handles requests to /abuse endpoint
func HandleAbuse(env *Env, w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return HTTPError(http.StatusBadRequest, "cant_read_request_body", err)
	}

	model := models.Abuse{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return HTTPError(http.StatusBadRequest, "cant_parse_request_body", err)
	}
	model.CreatedAt = time.Now().UTC().Format(time.RFC3339)

	err = env.DB.Write(models.DBAbuse, shortuuid.New(), model)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "cant_write_data", err)
	}

	sendEmail("abuse", model.From, map[string]string{
		"URLs": model.Urls,
		"Text": model.Text,
	})

	return nil
}

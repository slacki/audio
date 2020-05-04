package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/square/go-jose.v2/json"
)

// HandleInit handles application init requests
func HandleInit(env *Env, w http.ResponseWriter, r *http.Request) error {
	maintenance, err := strconv.ParseBool(os.Getenv("API_MAINTENANCE"))
	if err != nil {
		log.Println("Error parsing API_MAINTENANCE", err)
		maintenance = false
	}

	responseStruct := struct {
		Maintenance bool `json:"maintenance"`
	}{
		Maintenance: maintenance,
	}
	response, err := json.Marshal(responseStruct)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "", err)
	}
	fmt.Fprint(w, string(response))

	return nil
}

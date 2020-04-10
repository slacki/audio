package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"
)

// HandleUpload handles uploaded fles
func HandleUpload(env *Env, w http.ResponseWriter, r *http.Request) error {
	r.ParseMultipartForm(20000000) // 20M
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return HTTPError(http.StatusInternalServerError, "", err)
	}
	defer file.Close()

	// TODO: file validation: size, extension

	// some logging
	// TODO: move it further down, also log new filename
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// make sure uploads directory exists
	_ = os.Mkdir("./uploads", 0700)

	newfile := uuid.NewV4().String()
	f, err := os.OpenFile("./uploads/"+newfile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return HTTPError(http.StatusInternalServerError, "", err)

	}
	defer f.Close()
	io.Copy(f, file)

	responseStruct := struct {
		File string `json:"file"`
	}{
		File: newfile,
	}
	response, err := json.Marshal(responseStruct)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "", err)
	}

	fmt.Fprint(w, string(response))

	return nil
}

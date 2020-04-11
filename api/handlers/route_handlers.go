package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

const maxUploadSize = 15 * 1024 * 1024

// HandleUpload handles uploaded fles
func HandleUpload(env *Env, w http.ResponseWriter, r *http.Request) error {
	// validate file size
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		return HTTPError(http.StatusBadRequest, "file_too_big", err)
	}

	// parse and validate file and post parameters
	file, _, err := r.FormFile("file")
	if err != nil {
		return HTTPError(http.StatusBadRequest, "invalid_file", err)
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return HTTPError(http.StatusBadRequest, "invalid_file", err)
	}

	// check file type, DetectContentType only needs the first 512 bytes
	detectedFileType := http.DetectContentType(fileBytes)
	switch detectedFileType {
	case "audio/mpeg3", "audio/x-mpeg-3":
	case "audio/wav", "audio/x-wav":
	case "audio/ogg":
		break
	default:
		return HTTPError(http.StatusBadRequest, "invalid_file_type", err)
	}
	fileName := uuid.NewV4().String()
	fileEndings, err := mime.ExtensionsByType(detectedFileType)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "cant_read_file_type", err)
	}
	newPath := filepath.Join(env.UploadPath, fileName+fileEndings[0])
	log.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "cant_write_file", err)
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return HTTPError(http.StatusInternalServerError, "cant_write_file", err)
	}

	// write response
	responseStruct := struct {
		File string `json:"file"`
	}{
		File: newPath,
	}
	response, err := json.Marshal(responseStruct)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "", err)
	}
	fmt.Fprint(w, string(response))

	return nil
}

package handlers

import (
	models "api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gabriel-vasile/mimetype"
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
	file, fileHandler, err := r.FormFile("files")
	if err != nil {
		return HTTPError(http.StatusBadRequest, "invalid_file", err)
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return HTTPError(http.StatusBadRequest, "invalid_file", err)
	}

	// check mime type
	mimeFromRequest := http.DetectContentType(fileBytes)
	log.Println("Detected file type:", mimeFromRequest)
	switch mimeFromRequest {
	case "audio/mpeg3", "audio/x-mpeg-3", "audio/mpeg":
	case "audio/wav", "audio/x-wav", "audio/wave":
	case "audio/ogg", "application/ogg":
		break
	default:
		log.Println("Invalid file type:", mimeFromRequest)
		return HTTPError(http.StatusBadRequest, "invalid_file_type", err)
	}

	// get file extension and create new file path
	fileExtension := mimetype.Detect(fileBytes).Extension()
	if len(fileExtension) == 0 {
		return HTTPError(http.StatusInternalServerError, "cant_determine_file_ext", err)
	}
	fileName := uuid.NewV4().String()
	newPath := filepath.Join(env.UploadPath, fileName+fileExtension)

	// write file to the storage
	newFile, err := os.Create(newPath)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "cant_write_file", err)
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return HTTPError(http.StatusInternalServerError, "cant_write_file", err)
	}

	// store data
	env.DB.Write(models.DBAudio, fileName, models.Audio{
		File:         fileName + fileExtension,
		Hash:         fileName,
		Views:        0,
		OriginalName: fileHandler.Filename,
		CreatedAt:    time.RFC3339,
		ModifiedAt:   time.Now().String(),
	})

	// write response
	responseStruct := struct {
		Hash string `json:"hash"`
	}{
		Hash: fileName,
	}
	response, err := json.Marshal(responseStruct)
	if err != nil {
		return HTTPError(http.StatusInternalServerError, "", err)
	}
	fmt.Fprint(w, string(response))

	return nil
}

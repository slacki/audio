package handlers

import (
	models "api/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/lithammer/shortuuid/v3"
)

const maxUploadSize = 15 * 1024 * 1024

var allowedMimes = []string{
	"audio/mpeg3", "audio/x-mpeg-3", "audio/mpeg",
	"audio/wav", "audio/x-wav", "audio/wave",
	"audio/ogg", "application/ogg",
}

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
		return HTTPError(http.StatusBadRequest, "invalid_post_params", err)
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return HTTPError(http.StatusBadRequest, "cant_read_file", err)
	}

	// at this point we will write the file to the storage
	// then we can check mime types
	// get file extension and create new file path
	fileExtension := mimetype.Detect(fileBytes).Extension()
	if len(fileExtension) == 0 {
		return HTTPError(http.StatusInternalServerError, "cant_determine_file_ext", err)
	}
	fileName := shortuuid.New()
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

	// validate file by mime type
	mimeFromRequest := http.DetectContentType(fileBytes)
	mimeFromFile, err := getMimeTypeFromFile(newPath)
	if err != nil {
		log.Printf("Can't read mime type from file: %s", err.Error())
	}
	// if mime type determined form both sources is invalid
	if !allowedMimeType(mimeFromRequest) && !allowedMimeType(mimeFromFile) {
		// if this fails, we have to remove the file from the storage
		os.Remove(newPath)
		return HTTPError(http.StatusBadRequest, "invalid_file_type_mime", nil)
	}

	// store data
	err = env.DB.Write(models.DBAudio, fileName, models.Audio{
		File:         fileName + fileExtension,
		Hash:         fileName,
		Views:        0,
		OriginalName: fileHandler.Filename,
		RequestMIME:  mimeFromRequest,
		FileMIME:     mimeFromFile,
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
		ViewedAt:     time.Now().UTC().Format(time.RFC3339),
	})
	if err != nil {
		os.Remove(newPath)
		return HTTPError(http.StatusInternalServerError, "db_error", err)
	}

	// write response
	responseStruct := struct {
		Hash string `json:"hash"`
	}{
		Hash: fileName,
	}
	response, err := json.Marshal(responseStruct)
	if err != nil {
		env.DB.Delete(models.DBAudio, fileName)
		os.Remove(newPath)
		return HTTPError(http.StatusInternalServerError, "", err)
	}
	fmt.Fprint(w, string(response))

	return nil
}

func allowedMimeType(mime string) bool {
	for _, allowedMime := range allowedMimes {
		if mime == allowedMime {
			return true
		}
	}

	return false
}

func getMimeTypeFromFile(path string) (string, error) {
	out, err := execFileCmd(path)
	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(`.*:\s+(\w+\/[\w-]+)`)
	matches := r.FindAllStringSubmatch(out, -1)
	if len(matches) < 1 || len(matches[0]) < 2 {
		return "", errors.New("Could't parse `file` output: " + out)
	}

	return matches[0][1], nil
}

func execFileCmd(path string) (string, error) {
	cmd := exec.Command("file", "--mime-type", path)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

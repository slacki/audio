package handlers

import (
	"api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-acme/lego/log"
	"github.com/lithammer/shortuuid/v3"
	"gopkg.in/gomail.v2"
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

	// send email and ignore the errors
	sendEmail(model.From, model.Text)

	return nil
}

func sendEmail(from, text string) {
	host := os.Getenv("API_SMTP_HOST")
	user := os.Getenv("API_SMTP_USER")
	password := os.Getenv("API_SMTP_PASSWORD")
	sender := os.Getenv(("API_SMTP_SENDER_ADDRESS"))
	port, err := strconv.Atoi(os.Getenv("API_SMTP_PORT"))
	if err != nil {
		log.Println("Can't parse smtp port")
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", sender)
	m.SetHeader("Subject", "[shareaudio.cc] contact form")
	m.SetBody("text/plain", fmt.Sprintf("From: %s\n Message: %s\n", from, text))

	d := gomail.NewDialer(host, port, user, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println("Could't dial or send an email", err)
	}
}

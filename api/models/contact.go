package models

// DBContact is database name for Contact model
const DBContact = "contact"

// Contact represents the JSON data structure for contact
type Contact struct {
	Name      string `json:"name"`
	From      string `json:"from"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

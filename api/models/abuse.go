package models

// DBAbuse is database name for Abuse model
const DBAbuse = "abuse"

// Abuse represents the JSON data structure for abuse reposts
type Abuse struct {
	From      string `json:"from"`
	Urls      string `json:urls`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

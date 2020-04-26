package models

// DBAudio is database name for Audio model
const DBAudio = "audio"

// Audio represents the JSON data structure for audio files
type Audio struct {
	File         string `json:"file"`
	Hash         string `json:"hash"`
	Views        int    `json:"views"`
	OriginalName string `json:"original_name"`
	RequestMIME  string `json:"request_mime"`
	FileMIME     string `json:"file_mime"`
	CreatedAt    string `json:"created_at"`
	ViewedAt     string `json:"viewed_at"`
}

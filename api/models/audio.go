package models

const DBAudio = "audio"

type Audio struct {
	File         string `json:"file"`
	Hash         string `json:"hash"`
	Views        int    `json:"views"`
	OriginalName string `json:"original_name"`
	CreatedAt    string `json:"created_at"`
	ViewedAt     string `json:"viewed_at"`
}

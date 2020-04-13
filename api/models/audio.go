package models

const DBAudio = "audio"

type Audio struct {
	File         string `json:"file"`
	Hash         string `json:"hash"`
	Views        int    `json:"views"`
	OriginalName string `json:"original_name"`
}
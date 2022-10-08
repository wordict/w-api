package entity

import "time"

type AddVocabularyRequest struct {
	Word         string `json:"word"`
	UserID       string `json:"user_id"`
	FromLanguage string `json:"from_language"`
	ToLanguage   string `json:"to_language"`
}
type AddVocabularyResponse struct {
	ID           string    `json:"id"`
	Word         string    `json:"word"`
	FromLanguage string    `json:"from_language"`
	ToLanguage   string    `json:"to_language"`
	CreatedAt    time.Time `json:"created_at"`
	UserID       string    `json:"user_id"`
}
type Vocabulary struct {
	ID           string
	Word         string
	UserID       string
	FromLanguage string
	ToLanguage   string
	Translation  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type AddVocabularyRequestDTO struct {
	ID           string
	Word         string
	Translation  string
	UserID       string
	FromLanguage string
	ToLanguage   string
}

package entity

import "time"

type AddVocabularyRequest struct {
	Data string `json:"data"`
}
type AddVocabularyResponse struct {
	ID           string `json:"id"`
	Data         string `json:"data"`
	FromLanguage string `json:"from_language"`
	ToLanguage   string `json:"to_language"`
	CreatedAt    string `json:"created_at"`
	UserID       string `json:"user_id"`
}
type Vocabulary struct {
	ID           string
	Data         string
	UserID       string
	FromLanguage string
	ToLanguage   string
	Translation  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type AddVocabularyRequestDTO struct {
	ID           string
	Data         string
	Translation  string
	UserID       string
	FromLanguage string
	ToLanguage   string
}

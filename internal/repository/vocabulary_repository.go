package repository

import (
	"context"
	"github.com/wordict/w-api/internal/entity"
	"time"
)

func (r *Repository) AddVocabulary(ctx context.Context, dto entity.AddVocabularyRequestDTO) (entity.Vocabulary, error) {
	query := `INSERT INTO vocabulary (id, word, user_id, from_language, to_language, word_translation)
              VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING created_at;`
	row := r.db.QueryRowContext(ctx, query, dto.ID, dto.Word, dto.UserID, dto.FromLanguage, dto.ToLanguage, dto.Translation)

	var createdAt time.Time
	if err := row.Scan(&createdAt); err != nil {
		return entity.Vocabulary{}, err
	}
	return entity.Vocabulary{
		ID:           dto.ID,
		Word:         dto.Word,
		UserID:       dto.UserID,
		FromLanguage: dto.FromLanguage,
		ToLanguage:   dto.ToLanguage,
		Translation:  dto.Translation,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
	}, nil
}

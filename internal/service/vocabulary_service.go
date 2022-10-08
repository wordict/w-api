package service

import (
	"context"
	"github.com/rs/xid"
	"github.com/wordict/w-api/internal/entity"
	_ "github.com/wordict/w-api/internal/handler"
)

func (s *Service) AddVocabulary(ctx context.Context, request entity.AddVocabularyRequest) (entity.AddVocabularyResponse, error) {
	translation, err := s.client.Translate(request.Word, request.FromLanguage, request.ToLanguage)
	if err != nil {
		return entity.AddVocabularyResponse{}, err
	}
	dto := entity.AddVocabularyRequestDTO{
		ID:           xid.New().String(),
		Word:         request.Word,
		Translation:  translation,
		UserID:       request.UserID,
		FromLanguage: request.FromLanguage,
		ToLanguage:   request.ToLanguage,
	}

	vocabulary, err := s.repository.AddVocabulary(ctx, dto)
	return entity.AddVocabularyResponse{
		ID:           vocabulary.ID,
		Word:         vocabulary.Word,
		FromLanguage: vocabulary.FromLanguage,
		ToLanguage:   vocabulary.FromLanguage,
		CreatedAt:    vocabulary.CreatedAt,
		UserID:       vocabulary.UserID,
	}, nil
}

package repository

import (
	"context"
	"github.com/wordict/w-api/internal/entity"
)

func (r *Repository) CreateUser(ctx context.Context, dto entity.SignupRequestDTO) error {
	query := `
	INSERT INTO users (id, email, password)
	VALUES ($1,$2,$3)`
	row := r.db.QueryRowContext(ctx, query, dto.ID, dto.Email, dto.Password)
	return row.Err()
}

func (r *Repository) IsUserExist(ctx context.Context, dto entity.SignInRequestDTO) (entity.User, error) {
	query := `
	SELECT id, email, password FROM users 
	WHERE email=$1 AND password=$2`

	row := r.db.QueryRowContext(ctx, query, dto.Email, dto.Password)

	var user entity.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return entity.User{}, err
	}
	return user, nil
}

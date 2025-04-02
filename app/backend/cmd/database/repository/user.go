package repository

import (
	"context"
	"database/sql"
)

type User struct {
	ID string
}

type UserRepo interface {
	GetByID(ctx context.Context, userID string) (*User, error)
}

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{
		db: db,
	}
}

func (pr *PostgresUserRepo) GetByID(ctx context.Context, userID string) (*User, error) {
	return nil, nil
}

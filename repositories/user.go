package repositories

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nqtinh/go-gin-project/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, userReq *models.CreateUserReq) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
}

func newUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct{ db *sqlx.DB }

func (r *userRepository) CreateUser(ctx context.Context, userReq *models.CreateUserReq) error {
	query := `INSERT INTO "user" (username, password)  VALUES ($1, $2)`
	if _, err := Executor(ctx).Exec(query, userReq.Username, userReq.Password); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	q := `SELECT id, username, password FROM "user" WHERE username = $1`
	return user, Executor(ctx).Get(user, q, username)
}

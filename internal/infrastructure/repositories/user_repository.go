package persistence_repositories

import (
	domain "github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/core/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) GetAll() ([]domain.User, error) {
	users := []domain.User{}
	err := repository.db.Select(&users, "SELECT id, name FROM users")

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *UserRepository) Create(entity domain.User) error {
	_, err := repository.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", uuid.New(), entity.Name)

	if err != nil {
		return err
	}

	return nil
}

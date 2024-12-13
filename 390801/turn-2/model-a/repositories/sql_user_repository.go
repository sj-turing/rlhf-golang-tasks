// repositories/sql_user_repository.go
package repositories

import (
	"database/sql"

	"390801/turn-2/model-a/models"

	"github.com/pkg/errors"
)

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{db: db}
}

func (r *SQLUserRepository) Create(user *models.User) error {
	// implementation for SQL (PostgreSQL)
	_, err := r.db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", user.ID, user.Name)
	return err
}

func (r *SQLUserRepository) GetByID(id string) (*models.User, error) {
	// implementation for SQL
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}
	return user, nil
}

// Similar implementations for Update, Delete, and List

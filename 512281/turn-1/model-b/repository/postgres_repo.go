// postgres_repo.go
package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Implement UserRepository methods using PostgreSQL database operations
func (r *PostgresRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT id, username, email FROM users WHERE username = $1`
	var user models.User
	if err := r.db.GetContext(ctx, &user, query, username); err != nil {
		return nil, errors.Wrap(err, "failed to get user by username")
	}
	return &user, nil
}

// ... (Other database operations)

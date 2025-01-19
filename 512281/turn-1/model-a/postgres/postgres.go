package postgres

import (
	"512281/turn-1/model-a/models"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func (repo *PostgresUserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	row := repo.DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = $1", username)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Implement CreateUser, UpdateUser, DeleteUser similarly

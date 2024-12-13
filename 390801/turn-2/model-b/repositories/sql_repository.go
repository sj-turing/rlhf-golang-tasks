package repositories

import (
	"database/sql"
	"errors"
	"github.com/example/app/models"
)

type sqlUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) UserRepository {
	return &sqlUserRepository{db: db}
}

func (r *sqlUserRepository) Create(user *models.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users (id, name, email) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Name, user.Email)
	return err
}

func (r *sqlUserRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return &user, err
}

// ... similar implementations for Update, Delete, and List

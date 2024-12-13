package repositories

import (
	"errors"
	"sync"

	"390801/turn-1/model-b/models"
)

type inMemoryUserRepository struct {
	users []*models.User
	mu    sync.Mutex
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users: make([]*models.User, 0),
	}
}

func (r *inMemoryUserRepository) Create(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users = append(r.users, user)
	return nil
}

func (r *inMemoryUserRepository) GetByID(id int) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *inMemoryUserRepository) Update(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *inMemoryUserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *inMemoryUserRepository) List() ([]*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.users, nil
}

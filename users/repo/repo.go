package repo

import (
	"github.com/Disaxy/fast-track/users/entity"
	"github.com/Disaxy/fast-track/users/errors"
	"github.com/google/uuid"
	"log/slog"
)

type InMemoryUserRepo struct {
	users map[uuid.UUID]*entity.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[uuid.UUID]*entity.User),
	}
}

func (r *InMemoryUserRepo) Save(user *entity.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) FindByID(id uuid.UUID) (*entity.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.ErrUserNotFound
	}

	return user, nil
}

func (r *InMemoryUserRepo) FindAll() ([]*entity.User, error) {
	var users []*entity.User

	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *InMemoryUserRepo) DeleteByID(id uuid.UUID) error {
	delete(r.users, id)
	return nil
}

type MockUserRepo struct{}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{}
}

func (r *MockUserRepo) Save(user *entity.User) error {
	slog.Info("saving user")
	return nil
}

func (r *MockUserRepo) FindByID(id uuid.UUID) (*entity.User, error) {
	slog.Info("finding user")
	return nil, nil
}

func (r *MockUserRepo) FindAll() ([]*entity.User, error) {
	slog.Info("finding all users")
	return nil, nil
}

func (r *MockUserRepo) DeleteByID(id uuid.UUID) error {
	slog.Info("deleting user")
	return nil
}

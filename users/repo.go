package main

import "log/slog"

type UserRepository interface {
	Save(user User) error
	FindByID(id string) (User, error)
	FindAll() ([]User, error)
	DeleteByID(id string) error
}

type InMemoryUserRepo struct {
	users map[string]User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]User),
	}
}

func (r *InMemoryUserRepo) Save(user User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) FindByID(id string) (User, error) {
	user, ok := r.users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return user, nil
}

func (r *InMemoryUserRepo) FindAll() ([]User, error) {
	var users []User

	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

func (r *InMemoryUserRepo) DeleteByID(id string) error {
	delete(r.users, id)
	return nil
}

type MockUserRepo struct{}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{}
}

func (r *MockUserRepo) Save(user User) error {
	slog.Info("saving user")
	return nil
}

func (r *MockUserRepo) FindByID(id string) (User, error) {
	slog.Info("finding user")
	return User{}, nil
}

func (r *MockUserRepo) FindAll() ([]User, error) {
	slog.Info("finding all users")
	return []User{}, nil
}

func (r *MockUserRepo) DeleteByID(id string) error {
	slog.Info("deleting user")
	return nil
}

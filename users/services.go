package main

import "time"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(name, email, role string) (User, error) {
	user := User{
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}
	err := s.repo.Save(user)
	return user, err
}

func (s *UserService) GetUser(id string) (User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) ListUsers() []User {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil
	}
	return users
}

func (s *UserService) RemoveUser(id string) error {
	return s.repo.DeleteByID(id)
}

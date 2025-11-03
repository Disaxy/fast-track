package usecase

import (
	"github.com/Disaxy/fast-track/users/entity"
	"github.com/google/uuid"
	"time"
)

type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id uuid.UUID) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	DeleteByID(id uuid.UUID) error
}

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) UserUsecase {
	return UserUsecase{
		repo: repo,
	}
}

func (s *UserUsecase) CreateUser(name, email, role string) (*entity.User, error) {
	user := &entity.User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}
	err := s.repo.Save(user)
	return user, err
}

func (s *UserUsecase) GetUser(id uuid.UUID) (*entity.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserUsecase) ListUsers() []*entity.User {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil
	}
	return users
}

func (s *UserUsecase) RemoveUser(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}

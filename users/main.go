package main

import (
	"github.com/Disaxy/fast-track/users/repo"
	"github.com/Disaxy/fast-track/users/usecase"
	"log/slog"
	"os"
)

func main() {
	inMemoryRepo := repo.NewInMemoryUserRepo()
	userUsecase := usecase.NewUserUsecase(inMemoryRepo)

	newUser, err := userUsecase.CreateUser("Roma", "test@mail.com", "admin")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	user, err := userUsecase.GetUser(newUser.ID)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Create user", "user", user)
	slog.Info("List users", "users", userUsecase.ListUsers())

	err = userUsecase.RemoveUser(user.ID)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("User removed")
	slog.Info("List users", "users", userUsecase.ListUsers())
}

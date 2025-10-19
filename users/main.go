package main

import (
	"fmt"
	"os"
)

func main() {
	repo := NewInMemoryUserRepo()
	service := NewUserService(repo)

	newUser, err := service.CreateUser("Roma", "test@mail.com", "admin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	user, err := service.GetUser(newUser.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(user)
	fmt.Println(service.ListUsers())

	err = service.RemoveUser(user.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("User removed")
	fmt.Println(service.ListUsers())
}

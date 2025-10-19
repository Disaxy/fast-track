package main

const (
	ErrUserNotFound = Error("user not found")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

package main

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
}

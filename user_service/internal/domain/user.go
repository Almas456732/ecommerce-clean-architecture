package domain

import "time"

type User struct {
	ID        string
	Username  string
	Password  string
	Email     string
	FullName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

package user

import (
	"time"
)

type User struct {
	ID           int
	Nama         string
	Pekerjaan    string
	Email        string
	Password     string
	ImageProfile string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

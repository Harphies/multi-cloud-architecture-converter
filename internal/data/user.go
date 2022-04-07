package data

import (
	"database/sql"
	"errors"
	"time"
)

// custom password type
type password struct {
	plaintext *string
	hash      []byte
}

type Users struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
	Version   int       `json:"-"`
}

type UsersModel struct {
	DB *sql.DB
}

// define custom ErrDuplicateEmail error
var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

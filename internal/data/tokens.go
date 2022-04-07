package data

import (
	"database/sql"
	"time"
)

type Token struct {
	Plaintextt string    `json:"token"`
	Hash       []byte    `json:"-"`
	UserID     int64     `json:"-"`
	Expiry     time.Time `json:"expiry"`
	Scope      string    `json:"-"`
}

// token Model

type TokensModel struct {
	DB *sql.DB
}

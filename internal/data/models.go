package data

import (
	"database/sql"
	"errors"
)

// custom variables
var (
	ErrorNotFound    = errors.New("record not found")
	ErreEditConflict = errors.New("edit conflict")
)

// create a new Model
type Models struct {
	Movies      MovideModel
	Users       UsersModel
	Tokens      TokensModel
	Permissions PermissionsModel
}

func NewModels(db *sql.DB) Models {
	return Models{}
}

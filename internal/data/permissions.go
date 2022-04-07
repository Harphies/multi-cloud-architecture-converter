package data

import "database/sql"

type Permission []string

// permission model
type PermissionsModel struct {
	DB *sql.DB
}

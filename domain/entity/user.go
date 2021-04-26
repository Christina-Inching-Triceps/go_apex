package entity

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name sql.NullString
}

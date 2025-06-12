package entity

import (
	"database/sql"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Phone     string         `json:"phone" db:"phone"`
	Password  string         `json:"password,omitempty" db:"password"`
	RoleID    uuid.UUID      `json:"role_id" db:"role_id"`
	DeletedAt sql.NullTime   `json:"deleted_at" db:"deleted_at"`
	CreatedBy sql.NullString `json:"created_by" db:"created_by"`
	UpdatedBy sql.NullString `json:"updated_by" db:"updated_by"`
	DeletedBy sql.NullString `json:"deleted_by" db:"deleted_by"`
}

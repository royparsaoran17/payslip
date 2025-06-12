// Package entity
// Automatic generated
package entity

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Employee struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	UserID    string         `json:"user_id" db:"user_id"`
	Salary    float64        `json:"salary" db:"salary"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at" db:"deleted_at"`
	CreatedBy sql.NullString `json:"created_by" db:"created_by"`
	UpdatedBy sql.NullString `json:"updated_by" db:"updated_by"`
	DeletedBy sql.NullString `json:"deleted_by" db:"deleted_by"`
}

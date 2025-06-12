// Package entity
// Automatic generated
package entity

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Overtime struct {
	ID           uuid.UUID      `json:"id" db:"id"`
	EmployeeID   uuid.UUID      `json:"employee_id" db:"employee_id"`
	OvertimeDate time.Time      `json:"overtime_date" db:"overtime_date"`
	Hours        float64        `json:"hours" db:"hours"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt    sql.NullTime   `json:"deleted_at" db:"deleted_at"`
	CreatedBy    sql.NullString `json:"created_by" db:"created_by"`
	UpdatedBy    sql.NullString `json:"updated_by" db:"updated_by"`
	DeletedBy    sql.NullString `json:"deleted_by" db:"deleted_by"`
}

// Package entity
// Automatic generated
package entity

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Reimbursement struct {
	ID                uuid.UUID      `json:"id" db:"id"`
	EmployeeID        uuid.UUID      `json:"employee_id" db:"employee_id"`
	ReimbursementDate time.Time      `json:"reimbursement_date" db:"reimbursement_date"`
	Amount            float64        `json:"amount" db:"amount"`
	Description       string         `json:"description" db:"description"`
	Status            string         `json:"status" db:"status"` // e.g., pending, approved, rejected
	CreatedAt         time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt         sql.NullTime   `json:"deleted_at" db:"deleted_at"`
	CreatedBy         sql.NullString `json:"created_by" db:"created_by"`
	UpdatedBy         sql.NullString `json:"updated_by" db:"updated_by"`
	DeletedBy         sql.NullString `json:"deleted_by" db:"deleted_by"`
}

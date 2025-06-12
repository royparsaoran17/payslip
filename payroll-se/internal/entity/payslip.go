// Package entity
// Automatic generated
package entity

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Payslip struct {
	ID                 uuid.UUID      `json:"id" db:"id"`
	EmployeeID         uuid.UUID      `json:"employee_id" db:"employee_id"`
	PayrollPeriodID    uuid.UUID      `json:"payroll_period_id" db:"payroll_period_id"`
	BaseSalary         float64        `json:"base_salary" db:"base_salary"`
	ProratedSalary     float64        `json:"prorated_salary" db:"prorated_salary"`
	OvertimePay        float64        `json:"overtime_pay" db:"overtime_pay"`
	ReimbursementTotal float64        `json:"reimbursement_total" db:"reimbursement_total"`
	TakeHomePay        float64        `json:"take_home_pay" db:"take_home_pay"`
	CreatedAt          time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt          sql.NullTime   `json:"deleted_at" db:"deleted_at"`
	CreatedBy          sql.NullString `json:"created_by" db:"created_by"`
	UpdatedBy          sql.NullString `json:"updated_by" db:"updated_by"`
	DeletedBy          sql.NullString `json:"deleted_by" db:"deleted_by"`
}

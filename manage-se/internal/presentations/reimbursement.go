package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"
)

type ReimbursementCreate struct {
	EmployeeID        string    `json:"employee_id"`
	ReimbursementDate time.Time `json:"reimbursement_date"`
	Amount            float64   `json:"amount"`
	Description       string    `json:"description"`
	CreatedBy         string    `json:"created_by"`
}

func (r *ReimbursementCreate) Validate() error {
	return validation.Errors{
		"employee_id":        validation.Validate(&r.EmployeeID, validation.Required, is.UUID),
		"reimbursement_date": validation.Validate(&r.ReimbursementDate, validation.Required),
		"amount":             validation.Validate(&r.Amount, validation.Required),
		"description":        validation.Validate(&r.Description, validation.Required),
		"created_by":         validation.Validate(&r.CreatedBy, validation.Required),
	}.Filter()
}

type ReimbursementUpdate struct {
	ReimbursementDate time.Time `json:"reimbursement_date"`
	Amount            float64   `json:"amount"`
	Status            string    `json:"status"`
	UpdatedBy         string    `json:"updated_by"`
}

func (r *ReimbursementUpdate) Validate() error {
	return validation.Errors{
		"reimbursement_date": validation.Validate(&r.ReimbursementDate, validation.Required),
		"amount":             validation.Validate(&r.Amount, validation.Required),
		"status":             validation.Validate(&r.Status, validation.Required),
		"updated_by":         validation.Validate(&r.UpdatedBy, validation.Required),
	}.Filter()
}

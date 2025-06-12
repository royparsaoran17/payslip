package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type EmployeeCreate struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Salary    float64 `json:"salary"`
	CreatedBy string  `json:"created_by"`
}

func (e *EmployeeCreate) Validate() error {
	return validation.Errors{
		"user_id":    validation.Validate(&e.UserID, validation.Required, is.UUID),
		"salary":     validation.Validate(&e.Salary, validation.Required),
		"created_by": validation.Validate(&e.CreatedBy, validation.Required),
	}.Filter()
}

type EmployeeUpdate struct {
	Salary    float64 `json:"salary"`
	UpdatedBy string  `json:"updated_by"`
}

func (e *EmployeeUpdate) Validate() error {
	return validation.Errors{
		"salary":     validation.Validate(&e.Salary, validation.Required),
		"updated_by": validation.Validate(&e.UpdatedBy, validation.Required),
	}.Filter()
}

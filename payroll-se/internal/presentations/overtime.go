package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"
)

type OvertimeCreate struct {
	EmployeeID   string    `json:"employee_id"`
	OvertimeDate time.Time `json:"overtime_date"`
	Hours        float64   `json:"hours"`
	CreatedBy    string    `json:"created_by"`
}

func (o *OvertimeCreate) Validate() error {
	return validation.Errors{
		"employee_id":   validation.Validate(&o.EmployeeID, validation.Required, is.UUID),
		"overtime_date": validation.Validate(&o.OvertimeDate, validation.Required),
		"hours":         validation.Validate(&o.Hours, validation.Required),
		"created_by":    validation.Validate(&o.CreatedBy, validation.Required),
	}.Filter()
}

type OvertimeUpdate struct {
	OvertimeDate time.Time `json:"overtime_date"`
	Hours        float64   `json:"hours"`
	UpdatedBy    string    `json:"updated_by"`
}

func (o *OvertimeUpdate) Validate() error {
	return validation.Errors{
		"hours":         validation.Validate(&o.Hours, validation.Required),
		"overtime_date": validation.Validate(&o.OvertimeDate, validation.Required),
		"updated_by":    validation.Validate(&o.UpdatedBy, validation.Required),
	}.Filter()
}

package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"
)

type AttendanceCreate struct {
	EmployeeID     string    `json:"employee_id"`
	AttendanceDate time.Time `json:"attendance_date"`
	CreatedBy      string    `json:"created_by"`
}

func (a *AttendanceCreate) Validate() error {
	return validation.Errors{
		"employee_id":     validation.Validate(&a.EmployeeID, validation.Required, is.UUID),
		"attendance_date": validation.Validate(&a.AttendanceDate, validation.Required),
		"created_by":      validation.Validate(&a.CreatedBy, validation.Required),
	}.Filter()
}

type AttendanceUpdate struct {
	AttendanceDate time.Time `json:"attendance_date"`
	UpdatedBy      string    `json:"updated_by"`
}

func (a *AttendanceUpdate) Validate() error {
	return validation.Errors{
		"attendance_date": validation.Validate(&a.AttendanceDate, validation.Required),
		"updated_by":      validation.Validate(&a.UpdatedBy, validation.Required),
	}.Filter()
}

package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type PayslipCreate struct {
	EmployeeID         string  `json:"employee_id"`
	PayrollPeriodID    string  `json:"payroll_period_id"`
	BaseSalary         float64 `json:"base_salary"`
	ProratedSalary     float64 `json:"prorated_salary"`
	OvertimePay        float64 `json:"overtime_pay"`
	ReimbursementTotal float64 `json:"reimbursement_total"`
	TakeHomePay        float64 `json:"take_home_pay"`
	CreatedBy          string  `json:"created_by"`
}

func (p *PayslipCreate) Validate() error {
	return validation.Errors{
		"employee_id":         validation.Validate(&p.EmployeeID, validation.Required, is.UUID),
		"payroll_period_id":   validation.Validate(&p.PayrollPeriodID, validation.Required, is.UUID),
		"base_salary":         validation.Validate(&p.BaseSalary, validation.Required),
		"prorated_salary":     validation.Validate(&p.ProratedSalary, validation.Required),
		"overtime_pay":        validation.Validate(&p.OvertimePay, validation.Required),
		"reimbursement_total": validation.Validate(&p.ReimbursementTotal, validation.Required),
		"take_home_pay":       validation.Validate(&p.TakeHomePay, validation.Required),
		"created_by":          validation.Validate(&p.CreatedBy, validation.Required),
	}.Filter()
}

type PayslipUpdate struct {
	BaseSalary         float64 `json:"base_salary"`
	ProratedSalary     float64 `json:"prorated_salary"`
	OvertimePay        float64 `json:"overtime_pay"`
	ReimbursementTotal float64 `json:"reimbursement_total"`
	TakeHomePay        float64 `json:"take_home_pay"`
	UpdatedBy          string  `json:"updated_by"`
}

func (p *PayslipUpdate) Validate() error {
	return validation.Errors{
		"base_salary":         validation.Validate(&p.BaseSalary, validation.Required),
		"prorated_salary":     validation.Validate(&p.ProratedSalary, validation.Required),
		"overtime_pay":        validation.Validate(&p.OvertimePay, validation.Required),
		"reimbursement_total": validation.Validate(&p.ReimbursementTotal, validation.Required),
		"take_home_pay":       validation.Validate(&p.TakeHomePay, validation.Required),
		"updated_by":          validation.Validate(&p.UpdatedBy, validation.Required),
	}.Filter()
}

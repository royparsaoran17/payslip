package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type PayrollPeriodCreate struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedBy string    `json:"created_by"`
}

func (p *PayrollPeriodCreate) Validate() error {
	return validation.Errors{
		"start_date": validation.Validate(&p.StartDate, validation.Required),
		"end_date":   validation.Validate(&p.EndDate, validation.Required),
		"created_by": validation.Validate(&p.CreatedBy, validation.Required),
	}.Filter()
}

type PayrollPeriodUpdate struct {
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	UpdatedBy   string    `json:"updated_by"`
	IsProcessed bool      `json:"is_processed"`
}

func (p *PayrollPeriodUpdate) Validate() error {
	return validation.Errors{
		"start_date": validation.Validate(&p.StartDate, validation.Required),
		"end_date":   validation.Validate(&p.EndDate, validation.Required),
		"updated_by": validation.Validate(&p.UpdatedBy, validation.Required),
	}.Filter()
}

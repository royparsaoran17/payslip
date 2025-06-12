package payrollperiod

import (
	"payroll-se/pkg/databasex"
)

type payrollPeriod struct {
	db databasex.Adapter
}

func NewPayrollPeriod(db databasex.Adapter) *payrollPeriod {
	return &payrollPeriod{
		db: db,
	}
}

func (c *payrollPeriod) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *payrollPeriod) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}

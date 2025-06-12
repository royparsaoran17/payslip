package payslip

import (
	"payroll-se/pkg/databasex"
)

type payslip struct {
	db databasex.Adapter
}

func NewPayslip(db databasex.Adapter) *payslip {
	return &payslip{
		db: db,
	}
}

func (c *payslip) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *payslip) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}

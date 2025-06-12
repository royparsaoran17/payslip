package reimbursement

import (
	"payroll-se/pkg/databasex"
)

type reimbursement struct {
	db databasex.Adapter
}

func NewReimbursement(db databasex.Adapter) *reimbursement {
	return &reimbursement{
		db: db,
	}
}

func (c *reimbursement) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *reimbursement) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}

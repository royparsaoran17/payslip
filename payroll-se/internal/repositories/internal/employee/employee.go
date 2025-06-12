package employee

import (
	"payroll-se/pkg/databasex"
)

type employee struct {
	db databasex.Adapter
}

func NewEmployee(db databasex.Adapter) *employee {
	return &employee{
		db: db,
	}
}

func (c *employee) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *employee) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}

package attendance

import (
	"payroll-se/pkg/databasex"
)

type attendance struct {
	db databasex.Adapter
}

func NewAttendance(db databasex.Adapter) *attendance {
	return &attendance{
		db: db,
	}
}

func (c *attendance) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *attendance) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}

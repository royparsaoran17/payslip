package overtime

import (
	"payroll-se/pkg/databasex"
)

type overtime struct {
	db databasex.Adapter
}

func NewOvertime(db databasex.Adapter) *overtime {
	return &overtime{
		db: db,
	}
}

func (c *overtime) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at", "name":
		return true
	default:
		return false
	}

}

func (c *overtime) Searchable(field string) bool {
	switch field {
	case "name":
		return true
	default:
		return false
	}

}

package employee

import (
	"context"
	"github.com/pkg/errors"
	"payroll-se/internal/entity"
)

func (r employee) GetListEmployee(ctx context.Context) ([]entity.Employee, error) {
	query := `
		SELECT 
			id, 
			user_id, 
			salary,  
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz,
			created_by, 
			updated_by,
			deleted_by
		FROM employees 
			WHERE deleted_at is null
`
	employees := make([]entity.Employee, 0)

	err := r.db.Query(ctx, &employees, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all employees from database")
	}

	return employees, nil
}

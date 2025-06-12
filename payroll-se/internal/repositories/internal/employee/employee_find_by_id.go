package employee

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r employee) FindEmployeeByID(ctx context.Context, employeeID string) (*entity.Employee, error) {
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
		WHERE id = $1
		  AND deleted_at is null
`

	var employee entity.Employee

	err := r.db.QueryRow(ctx, &employee, query, employeeID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrDataNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &employee, nil
}

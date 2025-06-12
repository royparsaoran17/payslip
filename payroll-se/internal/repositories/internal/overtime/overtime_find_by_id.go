package overtime

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r overtime) FindOvertimeByID(ctx context.Context, overtimeID string) (*entity.Overtime, error) {
	query := `
		SELECT 
			id, ,
			employee_id, 
			overtime_date, 
			hours,
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz,
			created_by, 
			updated_by,
			deleted_by
		FROM overtimes 
		WHERE id = $1
		  AND deleted_at is null
`

	var overtime entity.Overtime

	err := r.db.QueryRow(ctx, &overtime, query, overtimeID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrDataNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &overtime, nil
}

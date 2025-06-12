package overtime

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r overtime) FindTotalOvertimeByEmployee(ctx context.Context, employeeID string, startDate, endDate time.Time) (*entity.WorkDay, error) {
	query := `
		SELECT COALESCE(SUM(hours), 0) as total_work_days FROM overtimes
WHERE employee_id = $1 AND overtime_date BETWEEN $2 AND $3
  AND deleted_at IS NULL;
`

	var overtime entity.WorkDay

	err := r.db.QueryRow(ctx, &overtime, query, employeeID, startDate, endDate)
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

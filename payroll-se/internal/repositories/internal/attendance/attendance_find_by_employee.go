package attendance

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r attendance) FindTotalAttendanceByEmployee(ctx context.Context, employeeID string, startDate, endDate time.Time) (*entity.WorkDay, error) {
	query := `
		SELECT COUNT(*) as total_work_days FROM attendances
WHERE employee_id = $1 AND attendance_date BETWEEN $2 AND $3
  AND deleted_at IS NULL;
`

	var attendance entity.WorkDay
	err := r.db.QueryRow(ctx, &attendance, query, employeeID, startDate, endDate)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrDataNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &attendance, nil
}

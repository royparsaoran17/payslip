package attendance

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r attendance) FindAttendanceByID(ctx context.Context, attendanceID string) (*entity.Attendance, error) {
	query := `
		SELECT 
			id, 
			employee_id, 
			attendance_date, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz,
			created_by, 
			updated_by,
			deleted_by
		FROM attendances 
		WHERE id = $1
		  AND deleted_at is null
`

	var attendance entity.Attendance

	err := r.db.QueryRow(ctx, &attendance, query, attendanceID)
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

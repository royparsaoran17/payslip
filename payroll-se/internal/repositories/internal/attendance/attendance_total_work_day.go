package attendance

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r attendance) TotalWorkDay(ctx context.Context, startDate, endDate time.Time) (*entity.WorkDay, error) {
	query := `
		SELECT COUNT(*) AS total_work_days
		FROM generate_series($1::date, $2::date, interval '1 day') d
		WHERE EXTRACT(dow FROM d) BETWEEN 1 AND 5;
`

	var attendance entity.WorkDay

	err := r.db.QueryRow(ctx, &attendance, query, startDate, endDate)
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

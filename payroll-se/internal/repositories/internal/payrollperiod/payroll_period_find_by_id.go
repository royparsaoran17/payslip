package payrollperiod

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r payrollPeriod) FindPayrollPeriodByID(ctx context.Context, payrollPeriodID string) (*entity.PayrollPeriod, error) {
	query := `
		SELECT 
			id, 
			start_date, 
			end_date, 
			is_processed, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz,
			created_by, 
			updated_by,
			deleted_by
		FROM payroll_periods 
		WHERE id = $1
		  AND deleted_at is null
`

	var payrollPeriod entity.PayrollPeriod

	err := r.db.QueryRow(ctx, &payrollPeriod, query, payrollPeriodID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrDataNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &payrollPeriod, nil
}

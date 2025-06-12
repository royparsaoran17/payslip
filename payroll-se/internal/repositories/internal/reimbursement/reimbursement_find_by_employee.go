package reimbursement

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r reimbursement) FindTotalReimbursementByEmployee(ctx context.Context, employeeID string, startDate, endDate time.Time) (*entity.WorkAmount, error) {
	query := `
		SELECT COALESCE(SUM(amount), 0) as total_amount FROM reimbursements
WHERE employee_id = $1 AND reimbursement_date BETWEEN $2 AND $3
  AND deleted_at IS NULL;
`

	var reimbursement entity.WorkAmount
	err := r.db.QueryRow(ctx, &reimbursement, query, employeeID, startDate, endDate)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrDataNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &reimbursement, nil
}

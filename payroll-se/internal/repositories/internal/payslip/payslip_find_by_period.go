package payslip

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r payslip) FindPayslipByPeriod(ctx context.Context, periodID string) (*entity.PayslipSummary, error) {
	query := `
		SELECT
		  COUNT(*) AS total_employees,
		  COALESCE(SUM(base_salary), 0) AS total_base_salary,
		  COALESCE(SUM(prorated_salary), 0) AS total_prorated_salary,
		  COALESCE(SUM(overtime_pay), 0) AS total_overtime,
		  COALESCE(SUM(reimbursement_total), 0) AS total_reimbursement,
		  COALESCE(SUM(take_home_pay), 0) AS total_take_home_pay
		FROM payslips
		WHERE payroll_period_id = $1
		  AND deleted_at IS NULL;
`

	var payslip entity.PayslipSummary

	err := r.db.QueryRow(ctx, &payslip, query, periodID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrDataNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &payslip, nil
}

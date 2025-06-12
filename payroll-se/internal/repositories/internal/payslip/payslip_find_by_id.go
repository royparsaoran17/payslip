package payslip

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r payslip) FindPayslipByID(ctx context.Context, payslipID string) (*entity.Payslip, error) {
	query := `
		SELECT 
			id, 
			employee_id, 
			payroll_period_id, 
			base_salary, 
			prorated_salary, 
			overtime_pay,
			reimbursement_total, 
			take_home_pay, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz,
			created_by, 
			updated_by,
			deleted_by
		FROM payslips 
		WHERE id = $1
		  AND deleted_at is null
`

	var payslip entity.Payslip

	err := r.db.QueryRow(ctx, &payslip, query, payslipID)
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

package payslip

import (
	"context"
	"github.com/pkg/errors"
	"payroll-se/internal/entity"
)

func (r payslip) GetAllByEmployee(ctx context.Context, employeeID, periodID string) ([]entity.Payslip, error) {

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
			WHERE employee_id = $1
			AND payroll_period_id = $2
			AND deleted_at IS NULL;
`

	payslips := make([]entity.Payslip, 0)

	err := r.db.Query(ctx, &payslips, query, employeeID, periodID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all payslips from database")
	}

	return payslips, nil
}

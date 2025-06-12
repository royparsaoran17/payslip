package payslip

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"payroll-se/internal/common"
	"payroll-se/internal/entity"
	"strings"
)

func (r payslip) GetAllPayslip(ctx context.Context, meta *common.Metadata) ([]entity.Payslip, error) {
	params, err := common.ParamFromMetadata(meta, &r)
	if err != nil {
		return nil, errors.Wrap(err, "parse params from meta")
	}

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
			WHERE 1=1
			AND deleted_at is null
			AND created_at >= GREATEST($3::date, '-infinity'::date)
			AND created_at <= LEAST($4::date, 'infinity'::date)
			ORDER BY created_at DESC
			LIMIT $1 OFFSET $2
`
	query = strings.Replace(
		query,
		"ORDER BY created_at DESC",
		fmt.Sprintf("ORDER BY %s %s", params.OrderBy, params.OrderDirection),
		1,
	)

	if params.SearchBy != "" {
		query = strings.Replace(
			query,
			"1=1",
			fmt.Sprintf("lower(%s) like '%s'", params.SearchBy, params.Search),
			1,
		)
	}

	payslips := make([]entity.Payslip, 0)

	err = r.db.Query(ctx, &payslips, query, params.Limit, params.Offset, params.DateFrom, params.DateEnd)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all payslips from database")
	}

	return payslips, nil
}

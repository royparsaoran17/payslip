package reimbursement

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
)

func (r reimbursement) FindReimbursementByID(ctx context.Context, reimbursementID string) (*entity.Reimbursement, error) {
	query := `
		SELECT 
			id, 
			employee_id, 
			reimbursement_date, 
			amount, 
			description, 
			status, 
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz,
			created_by, 
			updated_by,
			deleted_by
		FROM reimbursements 
		WHERE id = $1
		  AND deleted_at is null
`

	var reimbursement entity.Reimbursement

	err := r.db.QueryRow(ctx, &reimbursement, query, reimbursementID)
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

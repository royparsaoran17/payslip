package employee

import (
	"context"
	"database/sql"
	"payroll-se/internal/presentations"
	"payroll-se/internal/repositories/repooption"
	"payroll-se/pkg/postgres"
	"time"

	"github.com/pkg/errors"
	"payroll-se/internal/consts"
)

func (r employee) UpdateEmployee(ctx context.Context, employeeID string, input presentations.EmployeeUpdate, opts ...repooption.TxOption) error {

	txOpt := repooption.TxOptions{
		Tx:              nil,
		NotCommitInRepo: false,
	}

	for _, opt := range opts {
		opt(&txOpt)
	}

	if txOpt.Tx == nil {
		tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
		if err != nil {
			return errors.Wrap(err, "beginTx")
		}

		txOpt.Tx = tx
	}

	tx := txOpt.Tx
	query := `
	UPDATE employees 
	SET 
	    salary = $2, 
	    updated_at = $3,
	    updated_by = $4
	WHERE id = $1 
	AND deleted_at is null;`

	values := []interface{}{
		employeeID,
		input.Salary,
		time.Now().Local(),
		input.UpdatedBy,
	}

	if _, err := tx.ExecContext(ctx, query, values...); err != nil {
		if !txOpt.NotCommitInRepo {
			if err := tx.Rollback(); err != nil {
				err = errors.Wrap(err, "rollback")
			}
		}
		errSql := r.db.ParseSQLError(err)

		if errSql != nil {
			switch errSql {
			case postgres.ErrUniqueViolation:
				return consts.ErrDataAlreadyExist

			default:
				return errors.Wrap(err, "failed execute query")
			}
		}
	}

	if !txOpt.NotCommitInRepo {
		err := tx.Commit()
		if err != nil {
			return errors.Wrap(err, "commit add chopper")
		}
	}
	return nil
}

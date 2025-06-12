package overtime

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

func (r overtime) UpdateOvertime(ctx context.Context, overtimeID string, input presentations.OvertimeUpdate, opts ...repooption.TxOption) error {

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
	UPDATE overtimes 
	SET 
		overtime_date = $2, 
		hours = $3,
	    updated_at = $4,
	    updated_by = $5
	WHERE id = $1 
	AND deleted_at is null;`

	values := []interface{}{
		overtimeID,
		input.OvertimeDate,
		input.Hours,
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

package attendance

import (
	"context"
	"database/sql"
	"payroll-se/internal/repositories/repooption"
	"time"

	"payroll-se/internal/consts"
	"payroll-se/pkg/postgres"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"payroll-se/internal/presentations"
)

func (r attendance) CreateAttendance(ctx context.Context, input presentations.AttendanceCreate, opts ...repooption.TxOption) error {

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
	tNow := time.Now().Local()

	query := `
		INSERT INTO attendances (
			id, 
			employee_id, 
			attendance_date, 
			created_at, 
			updated_at, 
			created_by
		) VALUES ($1, $2, $3, $4, $4, $5)`

	values := []interface{}{
		uuid.New().String(),
		input.EmployeeID,
		input.AttendanceDate,
		tNow,
		input.CreatedBy,
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
				return consts.ErrDataNotFound

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

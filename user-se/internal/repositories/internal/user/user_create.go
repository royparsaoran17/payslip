package user

import (
	"auth-se/internal/common"
	"context"
	"database/sql"
	"time"

	"auth-se/internal/consts"
	"auth-se/pkg/postgres"

	"auth-se/internal/presentations"
	"github.com/pkg/errors"
)

func (c user) CreateUser(ctx context.Context, input presentations.UserCreate) error {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return errors.Wrap(err, "failed begin tx")
	}

	tNow := time.Now().Local()

	query := `INSERT INTO users (id, name, email, phone, password, role_id, created_at, updated_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $6, $7)`

	values := []interface{}{
		input.ID,
		input.Name,
		input.Email,
		input.Phone,
		common.Hasher(input.Password),
		input.RoleID,
		tNow,
		input.CreatedBy,
	}

	if _, err := c.db.ExecTx(ctx, tx, query, values...); err != nil {
		errRollback := c.db.RollbackTx(ctx, tx)
		if errRollback != nil {
			return errors.Wrap(err, "rollback failed")
		}

		errSql := c.db.ParseSQLError(err)

		if errSql != nil {
			switch errSql {
			case postgres.ErrUniqueViolation:
				return consts.ErrPhoneAlreadyExist

			default:
				return errors.Wrap(err, "failed execute query")
			}
		}

	}

	if err := c.db.CommitTx(ctx, tx); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

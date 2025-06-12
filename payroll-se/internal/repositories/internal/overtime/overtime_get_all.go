package overtime

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"payroll-se/internal/common"
	"payroll-se/internal/entity"
	"strings"
)

func (r overtime) GetAllOvertime(ctx context.Context, meta *common.Metadata) ([]entity.Overtime, error) {
	params, err := common.ParamFromMetadata(meta, &r)
	if err != nil {
		return nil, errors.Wrap(err, "parse params from meta")
	}

	query := `
		SELECT 
			id, ,
			employee_id, 
			overtime_date, 
			hours,
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz,
			created_by, 
			updated_by,
			deleted_by
		FROM overtimes 
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

	overtimes := make([]entity.Overtime, 0)

	err = r.db.Query(ctx, &overtimes, query, params.Limit, params.Offset, params.DateFrom, params.DateEnd)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all overtimes from database")
	}

	return overtimes, nil
}

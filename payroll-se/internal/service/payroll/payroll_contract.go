package payroll

import (
	"context"
	"payroll-se/internal/presentations"
)

type Payroll interface {
	SubmitAttendance(ctx context.Context, input presentations.AttendanceCreate) error
	SubmitOvertime(ctx context.Context, input presentations.OvertimeCreate) error
	SubmitReimbursement(ctx context.Context, input presentations.ReimbursementCreate) error
}

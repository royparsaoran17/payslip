package payroll

import (
	"context"
	"payroll-se/internal/entity"
	"payroll-se/internal/presentations"
)

type Payroll interface {
	CreatePayrollPeriod(ctx context.Context, input presentations.PayrollPeriodCreate) error
	SubmitAttendance(ctx context.Context, input presentations.AttendanceCreate) error
	SubmitOvertime(ctx context.Context, input presentations.OvertimeCreate) error
	SubmitReimbursement(ctx context.Context, input presentations.ReimbursementCreate) error
	GeneratePayslip(ctx context.Context, employeeID string, payrollPeriodID string) ([]entity.Payslip, error)
	GenerateSummary(ctx context.Context, payrollPeriodID string) (*entity.PayslipSummary, error)
	RunPayroll(ctx context.Context, input presentations.RunPayroll) error
}

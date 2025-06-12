package payroll

import (
	"context"
	"github.com/pkg/errors"
	"manage-se/internal/entity"
	"manage-se/internal/presentations"
	"manage-se/internal/provider"
)

type service struct {
	provider *provider.Provider
}

func NewService(provider *provider.Provider) Payroll {
	return &service{provider: provider}
}

func (s *service) SubmitAttendance(ctx context.Context, input presentations.AttendanceCreate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.provider.Payroll.CreateAttendance(ctx, input); err != nil {
		return errors.Wrap(err, "creating attendance")
	}

	return nil
}

func (s *service) SubmitOvertime(ctx context.Context, input presentations.OvertimeCreate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.provider.Payroll.CreateOvertime(ctx, input); err != nil {
		return errors.Wrap(err, "creating overtime")
	}

	return nil
}

func (s *service) CreatePayrollPeriod(ctx context.Context, input presentations.PayrollPeriodCreate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.provider.Payroll.CreatePayrollPeriod(ctx, input); err != nil {
		return errors.Wrap(err, "creating reimbursement")
	}

	return nil
}

func (s *service) SubmitReimbursement(ctx context.Context, input presentations.ReimbursementCreate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.provider.Payroll.CreateReimbursement(ctx, input); err != nil {
		return errors.Wrap(err, "creating reimbursement")
	}

	return nil
}

func (s *service) RunPayroll(ctx context.Context, input presentations.RunPayroll) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.provider.Payroll.RunPayroll(ctx, input); err != nil {
		return errors.Wrap(err, "running payroll")
	}

	return nil
}

func (s *service) GeneratePayslip(ctx context.Context, employeeID string, payrollPeriodID string) ([]entity.Payslip, error) {
	payslip, err := s.provider.Payroll.GetPayslipEmployee(ctx, employeeID, payrollPeriodID)
	if err != nil {
		return nil, errors.Wrap(err, "get employee attendance")
	}

	return payslip, nil
}

func (s *service) GenerateSummary(ctx context.Context, payrollPeriodID string) (*entity.PayslipSummary, error) {
	summary, err := s.provider.Payroll.GetPayslipSummary(ctx, payrollPeriodID)
	if err != nil {
		return nil, errors.Wrap(err, "get employee attendance")
	}

	return summary, nil
}

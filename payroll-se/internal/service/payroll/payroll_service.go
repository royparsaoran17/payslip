package payroll

import (
	"context"
	"github.com/pkg/errors"
	"payroll-se/internal/presentations"
	"payroll-se/internal/repositories"
)

type service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) Payroll {
	return &service{repo: repo}
}

func (s *service) SubmitAttendance(ctx context.Context, input presentations.AttendanceCreate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.repo.Attendance.CreateAttendance(ctx, input); err != nil {
		return errors.Wrap(err, "creating attendance")
	}

	return nil
}

func (s *service) SubmitOvertime(ctx context.Context, input presentations.OvertimeCreate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.repo.Overtime.CreateOvertime(ctx, input); err != nil {
		return errors.Wrap(err, "creating overtime")
	}

	return nil
}

func (s *service) SubmitReimbursement(ctx context.Context, input presentations.ReimbursementCreate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.repo.Reimbursement.CreateReimbursement(ctx, input); err != nil {
		return errors.Wrap(err, "creating reimbursement")
	}

	return nil
}

func (s *service) RunPayroll(ctx context.Context, userID string, input presentations.ReimbursementCreate) error {

	return nil
}

func (s *service) GeneratePayslip(ctx context.Context, userID string, input presentations.ReimbursementCreate) error {

	return nil
}

func (s *service) GenerateSummary(ctx context.Context, userID string, input presentations.ReimbursementCreate) error {

	return nil
}

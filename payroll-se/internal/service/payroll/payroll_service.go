package payroll

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"payroll-se/internal/consts"
	"payroll-se/internal/entity"
	"payroll-se/internal/presentations"
	"payroll-se/internal/repositories"
	"payroll-se/internal/repositories/repooption"
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

func (s *service) CreatePayrollPeriod(ctx context.Context, input presentations.PayrollPeriodCreate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.repo.PayrollPeriod.CreatePayrollPeriod(ctx, input); err != nil {
		return errors.Wrap(err, "creating reimbursement")
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

func (s *service) RunPayroll(ctx context.Context, input presentations.RunPayroll) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	period, err := s.repo.PayrollPeriod.FindPayrollPeriodByID(ctx, input.PayrollPeriodID)
	if err != nil {
		return errors.Wrap(err, "get payroll period")
	}

	employee, err := s.repo.Employee.GetListEmployee(ctx)
	if err != nil {
		return errors.Wrap(err, "get list of employee")
	}

	totalWorkDay, err := s.repo.Attendance.TotalWorkDay(ctx, period.StartDate, period.EndDate)
	if err != nil {
		return errors.Wrap(err, "get total work day")
	}

	// Start transaction
	tx, err := s.repo.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	for _, emp := range employee {
		var (
			workDays           = totalWorkDay.TotalWorkDays
			attendanceCount    int
			overtimeHours      int
			reimbursementTotal float64
			baseSalary         = emp.Salary
		)

		employeeAttendance, err := s.repo.Attendance.FindTotalAttendanceByEmployee(ctx, emp.ID.String(), period.StartDate, period.EndDate)
		if err != nil {
			_ = tx.Rollback()
			return errors.Wrap(err, "get employee attendance")
		}
		attendanceCount = employeeAttendance.TotalWorkDays

		employeeOvertime, err := s.repo.Overtime.FindTotalOvertimeByEmployee(ctx, emp.ID.String(), period.StartDate, period.EndDate)
		if err != nil {
			_ = tx.Rollback()
			return errors.Wrap(err, "get employee overtime")
		}
		overtimeHours = employeeOvertime.TotalWorkDays

		overtime, err := s.repo.Reimbursement.FindTotalReimbursementByEmployee(ctx, emp.ID.String(), period.StartDate, period.EndDate)
		if err != nil {
			_ = tx.Rollback()
			return errors.Wrap(err, "get employee attendance")
		}
		reimbursementTotal = overtime.TotalAmount

		var proratedSalary float64
		if workDays > 0 {
			proratedSalary = baseSalary * float64(attendanceCount) / float64(workDays)
		}

		overtimePay := float64(overtimeHours) * consts.OVERTIME_RATE_PER_HOUR
		takeHomePay := proratedSalary + overtimePay + reimbursementTotal

		if err = s.repo.Payslip.CreatePayslip(ctx, presentations.PayslipCreate{
			EmployeeID:         emp.ID.String(),
			PayrollPeriodID:    period.ID.String(),
			BaseSalary:         baseSalary,
			ProratedSalary:     proratedSalary,
			OvertimePay:        overtimePay,
			ReimbursementTotal: reimbursementTotal,
			TakeHomePay:        takeHomePay,
			CreatedBy:          input.CreatedBy,
		}, repooption.WithTx(tx)); err != nil {
			_ = tx.Rollback()
			return errors.Wrap(err, "creating reimbursement")
		}

	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

func (s *service) GeneratePayslip(ctx context.Context, employeeID string, payrollPeriodID string) ([]entity.Payslip, error) {
	payslip, err := s.repo.Payslip.GetAllByEmployee(ctx, employeeID, payrollPeriodID)
	if err != nil {
		return nil, errors.Wrap(err, "get employee attendance")
	}

	return payslip, nil
}

func (s *service) GenerateSummary(ctx context.Context, payrollPeriodID string) (*entity.PayslipSummary, error) {
	summary, err := s.repo.Payslip.FindPayslipByPeriod(ctx, payrollPeriodID)
	if err != nil {
		return nil, errors.Wrap(err, "get employee attendance")
	}

	return summary, nil
}

package repositories

import (
	"context"
	"payroll-se/internal/common"
	"payroll-se/internal/entity"
	"payroll-se/internal/presentations"
	"payroll-se/internal/repositories/repooption"
	"time"
)

type Overtime interface {
	CreateOvertime(ctx context.Context, input presentations.OvertimeCreate, opts ...repooption.TxOption) error
	UpdateOvertime(ctx context.Context, ID string, input presentations.OvertimeUpdate, opts ...repooption.TxOption) error
	FindOvertimeByID(ctx context.Context, ID string) (*entity.Overtime, error)
	GetAllOvertime(ctx context.Context, meta *common.Metadata) ([]entity.Overtime, error)
	FindTotalOvertimeByEmployee(ctx context.Context, employeeID string, startDate, endDate time.Time) (*entity.WorkDay, error)
}

type Payslip interface {
	CreatePayslip(ctx context.Context, input presentations.PayslipCreate, opts ...repooption.TxOption) error
	UpdatePayslip(ctx context.Context, ID string, input presentations.PayslipUpdate, opts ...repooption.TxOption) error
	FindPayslipByID(ctx context.Context, ID string) (*entity.Payslip, error)
	GetAllPayslip(ctx context.Context, meta *common.Metadata) ([]entity.Payslip, error)
	FindPayslipByPeriod(ctx context.Context, periodID string) (*entity.PayslipSummary, error)
	GetAllByEmployee(ctx context.Context, employeeID, periodID string) ([]entity.Payslip, error)
}

type Reimbursement interface {
	CreateReimbursement(ctx context.Context, input presentations.ReimbursementCreate, opts ...repooption.TxOption) error
	UpdateReimbursement(ctx context.Context, ID string, input presentations.ReimbursementUpdate, opts ...repooption.TxOption) error
	FindReimbursementByID(ctx context.Context, ID string) (*entity.Reimbursement, error)
	GetAllReimbursement(ctx context.Context, meta *common.Metadata) ([]entity.Reimbursement, error)
	FindTotalReimbursementByEmployee(ctx context.Context, employeeID string, startDate, endDate time.Time) (*entity.WorkAmount, error)
}

type PayrollPeriod interface {
	CreatePayrollPeriod(ctx context.Context, input presentations.PayrollPeriodCreate, opts ...repooption.TxOption) error
	UpdatePayrollPeriod(ctx context.Context, ID string, input presentations.PayrollPeriodUpdate, opts ...repooption.TxOption) error
	FindPayrollPeriodByID(ctx context.Context, ID string) (*entity.PayrollPeriod, error)
	GetAllPayrollPeriod(ctx context.Context, meta *common.Metadata) ([]entity.PayrollPeriod, error)
}

type Employee interface {
	CreateEmployee(ctx context.Context, input presentations.EmployeeCreate, opts ...repooption.TxOption) error
	UpdateEmployee(ctx context.Context, ID string, input presentations.EmployeeUpdate, opts ...repooption.TxOption) error
	FindEmployeeByID(ctx context.Context, ID string) (*entity.Employee, error)
	GetAllEmployee(ctx context.Context, meta *common.Metadata) ([]entity.Employee, error)
	GetListEmployee(ctx context.Context) ([]entity.Employee, error)
}

type Attendance interface {
	CreateAttendance(ctx context.Context, input presentations.AttendanceCreate, opts ...repooption.TxOption) error
	UpdateAttendance(ctx context.Context, ID string, input presentations.AttendanceUpdate, opts ...repooption.TxOption) error
	FindAttendanceByID(ctx context.Context, ID string) (*entity.Attendance, error)
	GetAllAttendance(ctx context.Context, meta *common.Metadata) ([]entity.Attendance, error)
	TotalWorkDay(ctx context.Context, startDate, endDate time.Time) (*entity.WorkDay, error)
	FindTotalAttendanceByEmployee(ctx context.Context, employeeID string, startDate, endDate time.Time) (*entity.WorkDay, error)
}

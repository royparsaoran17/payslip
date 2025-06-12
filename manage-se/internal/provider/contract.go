package provider

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/entity"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/user"
)

type User interface {
	Login(ctx context.Context, input presentations.Login) (*user.UserDetailToken, error)
	Verify(ctx context.Context, input presentations.Verify) (*user.UserDetail, error)

	CreateUser(ctx context.Context, input presentations.UserCreate) (*user.UserDetail, error)
	UpdateUser(ctx context.Context, userID string, input presentations.UserUpdate) (*user.UserDetail, error)
	GetListUsers(ctx context.Context, meta *common.Metadata) ([]user.User, error)
	GetUserByID(ctx context.Context, userID string) (*user.User, error)

	GetListRoles(ctx context.Context) ([]user.Role, error)
	CreateRole(ctx context.Context, input presentations.RoleCreate) (*user.Role, error)
	UpdateRole(ctx context.Context, roleID string, input presentations.RoleUpdate) (*user.Role, error)
}

type Payroll interface {
	CreateAttendance(ctx context.Context, input presentations.AttendanceCreate) error
	CreateOvertime(ctx context.Context, input presentations.OvertimeCreate) error
	CreatePayrollPeriod(ctx context.Context, input presentations.PayrollPeriodCreate) error
	CreateReimbursement(ctx context.Context, input presentations.ReimbursementCreate) error
	RunPayroll(ctx context.Context, input presentations.RunPayroll) error
	GetPayslipSummary(ctx context.Context, periodID string) (*entity.PayslipSummary, error)
	GetPayslipEmployee(ctx context.Context, employeeID, periodID string) ([]entity.Payslip, error)
}

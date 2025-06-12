package repositories

import (
	"context"
	"database/sql"
	"payroll-se/internal/repositories/internal/attendance"
	"payroll-se/internal/repositories/internal/employee"
	"payroll-se/internal/repositories/internal/overtime"
	"payroll-se/internal/repositories/internal/payrollperiod"
	"payroll-se/internal/repositories/internal/payslip"
	"payroll-se/internal/repositories/internal/reimbursement"
	"payroll-se/pkg/databasex"
)

type Repository struct {
	Attendance    Attendance
	Employee      Employee
	Overtime      Overtime
	Payslip       Payslip
	PayrollPeriod PayrollPeriod
	Reimbursement Reimbursement
	db            databasex.Adapter
}

func NewRepository(db databasex.Adapter) *Repository {
	return &Repository{
		Attendance:    attendance.NewAttendance(db),
		Employee:      employee.NewEmployee(db),
		Overtime:      overtime.NewOvertime(db),
		Payslip:       payslip.NewPayslip(db),
		PayrollPeriod: payrollperiod.NewPayrollPeriod(db),
		Reimbursement: reimbursement.NewReimbursement(db),
	}
}

func (r Repository) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, options)
}

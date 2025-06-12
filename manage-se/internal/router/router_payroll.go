package router

import (
	"manage-se/internal/handler"
	"manage-se/internal/ucase/payroll"
	"net/http"

	payrollsvc "manage-se/internal/service/payroll"
)

func (rtr *router) mountPayrolls(payrollSvc payrollsvc.Payroll) {
	rtr.router.HandleFunc("/external/v1/overtime", rtr.handle(
		handler.HttpRequest,
		payroll.NewOvertimeCreate(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/external/v1/payroll-period", rtr.handle(
		handler.HttpRequest,
		payroll.NewPayrollPeriodCreate(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/external/v1/reimbursement", rtr.handle(
		handler.HttpRequest,
		payroll.NewReimbursementCreate(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/external/v1/attendance", rtr.handle(
		handler.HttpRequest,
		payroll.NewAttendanceCreate(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/external/v1/run-payroll", rtr.handle(
		handler.HttpRequest,
		payroll.NewRunPayroll(payrollSvc),
	)).Methods(http.MethodPost)

	// ----

	rtr.router.HandleFunc("/external/v1/payslips/summary", rtr.handle(
		handler.HttpRequest,
		payroll.NewPayslipSummary(payrollSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/external/v1/employee/payslip", rtr.handle(
		handler.HttpRequest,
		payroll.NewPayslipEmployee(payrollSvc),
	)).Methods(http.MethodGet)

}

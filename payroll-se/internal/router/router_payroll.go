package router

import (
	"net/http"
	"payroll-se/internal/handler"
	"payroll-se/internal/ucase/payroll"

	payrollsvc "payroll-se/internal/service/payroll"
)

func (rtr *router) mountPayrolls(payrollSvc payrollsvc.Payroll) {
	rtr.router.HandleFunc("/internal/v1/overtime", rtr.handle(
		handler.HttpRequest,
		payroll.NewOvertimeCreate(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/payroll-period", rtr.handle(
		handler.HttpRequest,
		payroll.NewPayrollPeriodCreate(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/reimbursement", rtr.handle(
		handler.HttpRequest,
		payroll.NewReimbursementCreate(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/attendace", rtr.handle(
		handler.HttpRequest,
		payroll.NewAttendanceCreate(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/run-payroll", rtr.handle(
		handler.HttpRequest,
		payroll.NewRunPayroll(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/run-payroll", rtr.handle(
		handler.HttpRequest,
		payroll.NewRunPayroll(payrollSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/payslips/summary", rtr.handle(
		handler.HttpRequest,
		payroll.NewPayslipSummary(payrollSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/employee/{employee_id}/payslip", rtr.handle(
		handler.HttpRequest,
		payroll.NewPayslipEmployee(payrollSvc),
	)).Methods(http.MethodGet)

}

package payroll

import (
	"net/http"
	"payroll-se/internal/appctx"
	"payroll-se/internal/consts"
	"payroll-se/internal/service/payroll"
	"payroll-se/internal/ucase/contract"
	"payroll-se/pkg/tracer"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type payslipEmployee struct {
	service payroll.Payroll
}

func (r payslipEmployee) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.order_get_by_id")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("payslipEmployee")

	employeeID := mux.Vars(data.Request)["employee_id"]
	if _, err := uuid.Parse(employeeID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	periodID := data.Request.URL.Query().Get("period_id")

	result, err := r.service.GeneratePayslip(ctx, employeeID, periodID)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrOrderNotFound:
			return *responder.
				WithCode(http.StatusNotFound).
				WithMessage(causer.Error())

		default:
			tracer.SpanError(ctx, err)
			return *responder.
				WithCode(http.StatusInternalServerError).
				WithMessage(http.StatusText(http.StatusInternalServerError))
		}

	}

	return *responder.
		WithData(result).
		WithCode(http.StatusOK).
		WithMessage("order fetched")
}

func NewPayslipEmployee(service payroll.Payroll) contract.UseCase {
	return &payslipEmployee{service: service}
}

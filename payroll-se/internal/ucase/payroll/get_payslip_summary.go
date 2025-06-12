package payroll

import (
	"net/http"
	"payroll-se/internal/appctx"
	"payroll-se/internal/consts"
	"payroll-se/internal/service/payroll"
	"payroll-se/internal/ucase/contract"
	"payroll-se/pkg/tracer"

	"github.com/pkg/errors"
)

type payslipSummary struct {
	service payroll.Payroll
}

func (r payslipSummary) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.order_get_by_id")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("payslipSummary")

	periodID := data.Request.URL.Query().Get("period_id")

	result, err := r.service.GenerateSummary(ctx, periodID)
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

func NewPayslipSummary(service payroll.Payroll) contract.UseCase {
	return &payslipSummary{service: service}
}

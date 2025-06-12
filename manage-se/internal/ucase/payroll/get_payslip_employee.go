package payroll

import (
	"manage-se/internal/appctx"
	"manage-se/internal/consts"
	"manage-se/internal/entity"
	"manage-se/internal/service/payroll"
	"manage-se/internal/ucase/contract"
	"manage-se/pkg/logger"
	"manage-se/pkg/tracer"
	"net/http"

	"github.com/pkg/errors"
)

type payslipEmployee struct {
	service payroll.Payroll
}

func (r payslipEmployee) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.order_get_by_id")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("payslipEmployee")

	periodID := data.Request.URL.Query().Get("period_id")

	userCtx, ok := ctx.Value(consts.CtxUserAuth).(entity.UserContext)
	if !ok {
		logger.Error(errors.New("user data not exist in context"))
		return *responder.WithCode(http.StatusInternalServerError).WithMessage(http.StatusText(http.StatusInternalServerError))
	}

	result, err := r.service.GeneratePayslip(ctx, userCtx.ID, periodID)
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

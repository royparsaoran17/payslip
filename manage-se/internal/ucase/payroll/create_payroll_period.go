package payroll

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"manage-se/internal/consts"
	"manage-se/internal/entity"
	"manage-se/pkg/logger"
	"net/http"

	"manage-se/pkg/tracer"

	"github.com/pkg/errors"
	"manage-se/internal/appctx"
	"manage-se/internal/presentations"
	"manage-se/internal/service/payroll"
	"manage-se/internal/ucase/contract"
)

type payrollPeriodCreate struct {
	service payroll.Payroll
}

func (r payrollPeriodCreate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.payroll_period_create")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("payroll_period.Create")

	var input presentations.PayrollPeriodCreate
	if err := data.Cast(&input); err != nil {
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	userCtx, ok := ctx.Value(consts.CtxUserAuth).(entity.UserContext)
	if !ok {
		logger.Error(errors.New("user data not exist in context"))
		return *responder.WithCode(http.StatusInternalServerError).WithMessage(http.StatusText(http.StatusInternalServerError))
	}

	input.CreatedBy = userCtx.Email
	err := r.service.CreatePayrollPeriod(ctx, input)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		default:
			switch cause := causer.(type) {
			case consts.Error:
				return *responder.
					WithCode(http.StatusBadRequest).
					WithMessage(cause.Error())

			case validation.Errors:
				return *responder.
					WithCode(http.StatusUnprocessableEntity).
					WithError(cause).
					WithMessage("Validation error(s)")

			default:
				tracer.SpanError(ctx, err)
				return *responder.
					WithCode(http.StatusInternalServerError).
					WithMessage(http.StatusText(http.StatusInternalServerError))
			}
		}
	}

	return *responder.
		WithCode(http.StatusCreated).
		WithMessage("payroll period created")
}

func NewPayrollPeriodCreate(service payroll.Payroll) contract.UseCase {
	return &payrollPeriodCreate{service: service}
}

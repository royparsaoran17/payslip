package payroll

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"payroll-se/internal/consts"

	"payroll-se/pkg/tracer"

	"github.com/pkg/errors"
	"payroll-se/internal/appctx"
	"payroll-se/internal/presentations"
	"payroll-se/internal/service/payroll"
	"payroll-se/internal/ucase/contract"
)

type reimbursementCreate struct {
	service payroll.Payroll
}

func (r reimbursementCreate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.reimbursement_create")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("reimbursementCreate")

	var input presentations.ReimbursementCreate
	if err := data.Cast(&input); err != nil {
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	err := r.service.SubmitReimbursement(ctx, input)
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
		WithMessage("reimbursement created")
}

func NewReimbursementCreate(service payroll.Payroll) contract.UseCase {
	return &reimbursementCreate{service: service}
}

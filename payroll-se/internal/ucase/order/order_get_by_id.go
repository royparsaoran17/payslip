package order

import (
	"net/http"

	"payroll-se/pkg/tracer"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"payroll-se/internal/appctx"
	"payroll-se/internal/consts"
	"payroll-se/internal/service/order"
	"payroll-se/internal/ucase/contract"
)

type orderGetByID struct {
	service order.Order
}

func (r orderGetByID) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.order_get_by_id")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("orderGetByID")

	orderID := mux.Vars(data.Request)["order_id"]
	if _, err := uuid.Parse(orderID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	result, err := r.service.GetOrderByID(ctx, orderID)
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

func NewOrderGetByID(service order.Order) contract.UseCase {
	return &orderGetByID{service: service}
}

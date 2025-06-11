package user

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"manage-se/internal/entity"
	"manage-se/internal/provider/providererrors"
	"manage-se/internal/service/user"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"manage-se/internal/appctx"
	"manage-se/internal/consts"
	"manage-se/internal/presentations"
	"manage-se/internal/ucase/contract"
	"manage-se/pkg/logger"
	"manage-se/pkg/tracer"
)

type userUpdate struct {
	service user.User
}

func (c userUpdate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.user_update")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("userUpdate")
	var input presentations.UserUpdate

	if err := data.Cast(&input); err != nil {
		logger.Warn(fmt.Sprintf("error cast to userUpdate presentation %+v", err))
		tracer.SpanError(ctx, err)
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	userID := mux.Vars(data.Request)["user_id"]
	if _, err := uuid.Parse(userID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	userCtx, ok := ctx.Value(consts.CtxUserAuth).(entity.UserContext)
	if !ok {
		logger.Error(errors.New("user data not exist in context"))
		return *responder.WithCode(http.StatusInternalServerError).WithMessage(http.StatusText(http.StatusInternalServerError))
	}

	input.UpdatedBy = userCtx.Email
	err := c.service.UpdateUser(ctx, userID, input)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case providererrors.ErrUserNotFound:
			return *responder.
				WithCode(http.StatusNotFound).
				WithMessage(causer.Error())

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
		WithCode(http.StatusOK).
		WithMessage("user updated")
}

func NewUserUpdate(service user.User) contract.UseCase {
	return &userUpdate{service: service}
}

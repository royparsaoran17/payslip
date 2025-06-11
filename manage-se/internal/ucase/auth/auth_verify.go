package auth

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"manage-se/internal/consts"
	"manage-se/internal/entity"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/providererrors"
	"manage-se/internal/service/auth"
	"manage-se/pkg/httpx"
	"manage-se/pkg/jwt"
	"net/http"

	"manage-se/pkg/tracer"

	"github.com/pkg/errors"
	"manage-se/internal/appctx"
	"manage-se/internal/ucase/contract"
)

type verify struct {
	service auth.Auth
}

func (c verify) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.verify")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("verify")

	bearerToken := entity.BearerToken(data.Request.Header.Get(httpx.Authorization))

	if bearerToken.TokenEmpty() {
		return *responder.
			WithCode(http.StatusUnauthorized).
			WithMessage(http.StatusText(http.StatusUnauthorized))
	}

	users, err := c.service.VerifyToken(ctx, presentations.Verify{
		Token: bearerToken.GetToken(),
	})
	if err != nil {
		errCause := errors.Cause(err)
		switch errCause {
		case jwt.ErrTokenExpired:
			return *responder.
				WithCode(http.StatusUnauthorized).
				WithMessage(errCause.Error())

		default:
			switch causer := errCause.(type) {
			case consts.Error:
				return *responder.WithContext(ctx).WithCode(http.StatusBadRequest).WithMessage(errCause.Error())

			case providererrors.Error:
				return *responder.WithContext(ctx).WithCode(causer.Code).WithError(causer.Errors).WithMessage(causer.Message)

			case validation.Errors:
				return *responder.
					WithContext(ctx).
					WithCode(http.StatusUnprocessableEntity).
					WithMessage("Validation Error(s)").
					WithError(errCause)

			default:
				return *responder.WithContext(ctx).WithCode(http.StatusInternalServerError).WithMessage(http.StatusText(http.StatusInternalServerError))
			}
		}
	}

	return *responder.
		WithCode(http.StatusOK).
		WithData(users).
		WithMessage("user verified")
}

func NewVerify(service auth.Auth) contract.UseCase {
	return &verify{service: service}
}

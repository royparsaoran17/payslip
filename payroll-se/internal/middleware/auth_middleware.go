package middleware

import (
	"net/http"
	"payroll-se/internal/appctx"
)

func Authorize() MiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request, conf *appctx.Config) error {
		return nil
	}
}

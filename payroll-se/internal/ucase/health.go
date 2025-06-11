package ucase

import (
	"payroll-se/internal/appctx"
	"payroll-se/internal/consts"
	"payroll-se/internal/ucase/contract"
)

type healthCheck struct {
}

func NewHealthCheck() contract.UseCase {
	return &healthCheck{}
}

func (u *healthCheck) Serve(*appctx.Data) appctx.Response {
	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("ok")
}

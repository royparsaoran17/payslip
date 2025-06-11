package router

import (
	"manage-se/internal/consts"
	"manage-se/internal/middleware"
	authsvc "manage-se/internal/service/auth"
	"manage-se/internal/ucase/user"
	"net/http"

	"manage-se/internal/handler"
	usersvc "manage-se/internal/service/user"
)

func (rtr *router) mountUser(userSvc usersvc.User, authSvc authsvc.Auth) {
	rtr.router.HandleFunc("/external/v1/users", rtr.handle(
		handler.HttpRequest,
		user.NewUserGetAll(userSvc),
		middleware.Authorize(authSvc, consts.SuperAdminRoleName),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/external/v1/users/{user_id}", rtr.handle(
		handler.HttpRequest,
		user.NewUserGetByID(userSvc),
		middleware.Authorize(authSvc, consts.SuperAdminRoleName),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/external/v1/users", rtr.handle(
		handler.HttpRequest,
		user.NewUserCreate(userSvc),
		middleware.Authorize(authSvc, consts.SuperAdminRoleName),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/external/v1/users/{user_id}", rtr.handle(
		handler.HttpRequest,
		user.NewUserUpdate(userSvc),
		middleware.Authorize(authSvc, consts.SuperAdminRoleName),
	)).Methods(http.MethodPut)

}

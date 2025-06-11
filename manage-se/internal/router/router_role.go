package router

import (
	"manage-se/internal/ucase/role"
	"net/http"

	"manage-se/internal/handler"
	rolesvc "manage-se/internal/service/role"
)

func (rtr *router) mountRole(roleSvc rolesvc.Role) {
	rtr.router.HandleFunc("/external/v1/roles", rtr.handle(
		handler.HttpRequest,
		role.NewRoleGetAll(roleSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/external/v1/roles", rtr.handle(
		handler.HttpRequest,
		role.NewRoleCreate(roleSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/external/v1/roles/{role_id}", rtr.handle(
		handler.HttpRequest,
		role.NewRoleUpdate(roleSvc),
	)).Methods(http.MethodPut)
}

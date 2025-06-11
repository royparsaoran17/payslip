package role

import (
	"context"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/user"
)

type Role interface {
	GetAllRole(ctx context.Context) ([]user.Role, error)
	UpdateRole(ctx context.Context, roleID string, input presentations.RoleUpdate) error
	CreateRole(ctx context.Context, input presentations.RoleCreate) error
}

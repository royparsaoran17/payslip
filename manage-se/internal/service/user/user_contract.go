package user

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/user"
)

type User interface {
	GetAllUser(ctx context.Context, meta *common.Metadata) ([]user.User, error)
	GetUsetByID(ctx context.Context, userID string) (*user.User, error)
	UpdateUser(ctx context.Context, userID string, input presentations.UserUpdate) error
	CreateUser(ctx context.Context, input presentations.UserCreate) error
}

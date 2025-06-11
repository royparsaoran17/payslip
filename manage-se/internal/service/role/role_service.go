package role

import (
	"context"
	"manage-se/internal/presentations"
	"manage-se/internal/provider"
	"manage-se/internal/provider/user"

	"github.com/pkg/errors"
)

type service struct {
	provider *provider.Provider
}

func NewService(provider *provider.Provider) Role {
	return &service{provider: provider}
}

func (s *service) GetAllRole(ctx context.Context) ([]user.Role, error) {
	roles, err := s.provider.User.GetListRoles(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting all roles on ")
	}

	return roles, nil
}

func (s *service) UpdateRole(ctx context.Context, roleID string, input presentations.RoleUpdate) error {
	_, err := s.provider.User.UpdateRole(ctx, roleID, input)
	if err != nil {
		return nil
	}

	return nil
}

func (s *service) CreateRole(ctx context.Context, input presentations.RoleCreate) error {
	_, err := s.provider.User.CreateRole(ctx, input)
	if err != nil {
		return nil
	}

	return nil
}

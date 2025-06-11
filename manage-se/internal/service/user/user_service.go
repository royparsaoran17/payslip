package user

import (
	"context"
	"manage-se/internal/common"
	"manage-se/internal/presentations"
	"manage-se/internal/provider"
	"manage-se/internal/provider/user"

	"github.com/pkg/errors"
)

type service struct {
	provider *provider.Provider
}

func NewService(provider *provider.Provider) User {
	return &service{provider: provider}
}

func (s *service) GetAllUser(ctx context.Context, meta *common.Metadata) ([]user.User, error) {
	users, err := s.provider.User.GetListUsers(ctx, meta)
	if err != nil {
		return nil, errors.Wrap(err, "getting all users ")
	}

	return users, nil
}

func (s *service) GetUsetByID(ctx context.Context, userID string) (*user.User, error) {
	users, err := s.provider.User.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "getting all users ")
	}

	return users, nil
}

func (s *service) UpdateUser(ctx context.Context, userID string, input presentations.UserUpdate) error {
	_, err := s.provider.User.UpdateUser(ctx, userID, input)
	if err != nil {
		return nil
	}

	return nil
}

func (s *service) CreateUser(ctx context.Context, input presentations.UserCreate) error {
	_, err := s.provider.User.CreateUser(ctx, input)
	if err != nil {
		return nil
	}

	return nil
}

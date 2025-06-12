package employee

import (
	"context"
	"github.com/pkg/errors"
	"payroll-se/internal/entity"
	"payroll-se/internal/presentations"
	"payroll-se/internal/repositories"
)

type service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) Employee {
	return &service{repo: repo}
}

func (s *service) CreateEmployee(ctx context.Context, input presentations.EmployeeCreate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.repo.Employee.CreateEmployee(ctx, input); err != nil {
		return errors.Wrap(err, "creating attendance")
	}

	return nil
}

func (s *service) UpdateEmployee(ctx context.Context, ID string, input presentations.EmployeeUpdate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	if err := s.repo.Employee.UpdateEmployee(ctx, ID, input); err != nil {
		return errors.Wrap(err, "creating attendance")
	}

	return nil
}

func (s *service) GetEmployee(ctx context.Context, ID string) (*entity.Employee, error) {
	employeeData, err := s.repo.Employee.FindEmployeeByID(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "creating attendance")
	}

	return employeeData, nil
}

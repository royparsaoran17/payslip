package employee

import (
	"context"
	"payroll-se/internal/entity"
	"payroll-se/internal/presentations"
)

type Employee interface {
	CreateEmployee(ctx context.Context, input presentations.EmployeeCreate) error
	UpdateEmployee(ctx context.Context, ID string, input presentations.EmployeeUpdate) error
	GetEmployee(ctx context.Context, ID string) (*entity.Employee, error)
}

package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserUpdate struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	RoleID    string `json:"role_id"`
	UpdatedBy string `json:"updated_by"`
}

func (r *UserUpdate) Validate() error {
	return validation.Errors{
		"name":       validation.Validate(&r.Name, validation.Required),
		"phone":      validation.Validate(&r.Phone, validation.Required, is.E164),
		"role_id":    validation.Validate(&r.RoleID, validation.Required, is.UUID),
		"updated_by": validation.Validate(&r.UpdatedBy, validation.Required),
	}.Filter()
}

type UserCreate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	RoleID    string `json:"role_id"`
	CreatedBy string `json:"created_by"`
}

func (r *UserCreate) Validate() error {
	return validation.Errors{
		"name":       validation.Validate(&r.Name, validation.Required),
		"phone":      validation.Validate(&r.Phone, validation.Required, is.E164),
		"password":   validation.Validate(&r.Password, validation.Required),
		"role_id":    validation.Validate(&r.RoleID, validation.Required, is.UUID),
		"created_by": validation.Validate(&r.CreatedBy, validation.Required),
	}.Filter()
}

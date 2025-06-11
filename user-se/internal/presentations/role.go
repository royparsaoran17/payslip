package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RoleUpdate struct {
	Name      string `json:"name"`
	UpdatedBy string `json:"updated_by"`
}

func (r *RoleUpdate) Validate() error {
	return validation.Errors{
		"name":       validation.Validate(&r.Name, validation.Required),
		"updated_by": validation.Validate(&r.UpdatedBy, validation.Required),
	}.Filter()
}

type RoleCreate struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}

func (r *RoleCreate) Validate() error {
	return validation.Errors{
		"name":       validation.Validate(&r.Name, validation.Required),
		"created_by": validation.Validate(&r.CreatedBy, validation.Required),
	}.Filter()
}

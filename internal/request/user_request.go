package request

import (
	"be-park-ease/constants"
	"be-park-ease/internal/sql"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AllUserRequest struct {
	BasePagination
	Role   sql.UserRole
	Status sql.UserStatus
}

func (r *AllUserRequest) Normalize() {
	if !r.Role.Valid() {
		r.Role = ""
	}

	if !r.Status.Valid() {
		r.Status = ""
	}

	r.BasePagination.Normalize()
}

type CreateUserRequest struct {
	Username string         `json:"username" example:"paijo"`
	Password string         `json:"password" example:"password"`
	Name     string         `json:"name" example:"Paijo Royo Royo"`
	Role     sql.UserRole   `json:"role" example:"karyawan"`
	Status   sql.UserStatus `json:"status" example:"active"`
}

func (req CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(5, 100)),
		validation.Field(&req.Username, validation.Required, constants.ValidationUsername),
		validation.Field(&req.Password, validation.Required, constants.ValidationPassword),
		validation.Field(&req.Role, validation.Required, validation.In(sql.UserRoleAdmin, sql.UserRoleKaryawan)),
		validation.Field(&req.Status, validation.Required, validation.In(sql.UserStatusActive, sql.UserStatusBanned)),
	)
}

type UpdateUserRequest struct {
	CreateUserRequest
	UserId int32 `json:"-"`
}

func (req UpdateUserRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.When(req.Name != "", validation.Required, validation.Length(5, 100))),
		validation.Field(&req.Username, validation.When(req.Username != "", validation.Required, constants.ValidationUsername)),
		validation.Field(&req.Password, validation.When(req.Password != "", validation.Required, constants.ValidationPassword)),
		validation.Field(&req.Role, validation.When(req.Role != "", validation.Required, validation.In(sql.UserRoleAdmin, sql.UserRoleKaryawan))),
		validation.Field(&req.Status, validation.When(req.Status != "", validation.Required, validation.In(sql.UserStatusActive, sql.UserStatusBanned))),
	)
}

type ChangePasswordRequest struct {
	OldPassword          string `json:"old_password" example:"Password123@"`
	Password             string `json:"password" example:"Password123@"`
	PasswordConfirmation string `json:"password_confirmation" example:"Password123@"`
	UserId               int32  `json:"-"`
}

func (req ChangePasswordRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.OldPassword, validation.Required),
		validation.Field(&req.Password, validation.Required, constants.ValidationPassword),
		validation.Field(&req.PasswordConfirmation, validation.Required, validation.In(req.Password).Error("the password not equal with Password")),
	)
}

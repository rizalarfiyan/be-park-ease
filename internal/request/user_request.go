package request

import (
	"be-park-ease/internal/sql"
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

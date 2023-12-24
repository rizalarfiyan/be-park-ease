package response

import "be-park-ease/internal/sql"

type User struct {
	ID       int32          `json:"id"`
	Username string         `json:"username"`
	Name     string         `json:"name"`
	Role     sql.UserRole   `json:"role"`
	Status   sql.UserStatus `json:"status"`
}

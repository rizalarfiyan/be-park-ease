package response

import (
	"be-park-ease/internal/sql"
	"time"
)

type AuthLoginResponse struct {
	ID        int32          `json:"id"`
	Name      string         `json:"name"`
	Username  string         `json:"username"`
	Role      sql.UserRole   `json:"role"`
	Status    sql.UserStatus `json:"status"`
	Token     string         `json:"token"`
	ExpiredAt time.Time      `json:"expired_at"`
}

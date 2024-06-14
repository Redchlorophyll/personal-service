package response

import "time"

type GetAccountRespositoryResponse struct {
	SessionTokenExpiredAt *time.Time `db:"last_logged_in"`
	SessionToken          *string    `db:"session_token"`
	Username              string     `db:"user_name"`
	Email                 string     `db:"email"`
	FullName              string     `db:"full_name"`
	Description           *string    `db:"description"`
	Password              string     `db:"password"`
}

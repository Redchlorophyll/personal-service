package request

import "mime/multipart"

type CreateAccountRequest struct {
	UserName    string                `json:"user_name"`
	FullName    string                `json:"full_name"`
	ProfileImg  *multipart.FileHeader `json:"profile_img"`
	Description string                `json:"description"`
	Email       string                `json:"email"`
	Password    string                `json:"password"`
}

package request

type ProfileData struct {
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	FullName    string  `json:"full_name"`
	Description *string `json:"description"`
}

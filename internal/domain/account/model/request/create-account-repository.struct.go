package request

type CreateAccountRepositoryRequest struct {
	ProfileImgUrl string
	UserName      string
	Description   string
	FullName      string
	Password      string
	Email         string
}

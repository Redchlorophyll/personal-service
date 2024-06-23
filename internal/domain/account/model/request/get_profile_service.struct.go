package request

type GetProfileServiceRequest struct {
	// if Id is provided, will bypass Session token check.
	Id           *int64
	SessionToken string
}

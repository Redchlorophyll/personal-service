package request

import "errors"

type UpdateLinkyRequest struct {
	CreateLinkyRequest

	Id int
}

func (request *UpdateLinkyRequest) Validator() error {
	if request.Id == 0 {
		return errors.New("invalid ContentIdentiferId")
	}

	return request.CreateLinkyRequest.Validator()
}

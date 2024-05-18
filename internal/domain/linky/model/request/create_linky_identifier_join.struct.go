package request

import "errors"

type CreateLinkyIdentifierRequest struct {
	ContentIdentiferId *int
	LinkId             *int
}

func (request *CreateLinkyIdentifierRequest) Validator() error {
	if request.ContentIdentiferId == nil {
		return errors.New("invalid ContentIdentiferId")
	}

	if request.LinkId == nil {
		return errors.New("invalid LinkId")
	}

	return nil
}

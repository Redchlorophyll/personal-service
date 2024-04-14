package request

import "errors"

type CreateIdentifierRequest struct {
	Identifier string `json:"identifier"`
	Title      string `json:"title"`
}

func (request *CreateIdentifierRequest) Validator() error {
	if request.Identifier == "" {
		return errors.New("invalid identifier")
	}

	if request.Title == "" {
		return errors.New("invalid title")
	}

	return nil
}

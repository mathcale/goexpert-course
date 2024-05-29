package dto

import "errors"

type InputUCInput struct {
	Zipcode string `json:"cep"`
}

func (i InputUCInput) Validate() error {
	if i.Zipcode == "" || len(i.Zipcode) != 8 {
		return errors.New("invalid zipcode")
	}
	return nil
}

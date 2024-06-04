package dto

import (
	"errors"

	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/customerrors"
)

type InputUCInput struct {
	Zipcode string `json:"cep"`
}

func (i InputUCInput) Validate() error {
	if i.Zipcode == "" || len(i.Zipcode) != 8 {
		return &customerrors.ValidationError{
			Err:     errors.New("invalid zipcode"),
			Message: "invalid zipcode",
			Reasons: []string{"zipcode must have 8 characters"},
		}
	}

	return nil
}

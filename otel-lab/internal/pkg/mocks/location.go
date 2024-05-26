package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/mathcale/goexpert-course/otel-lab/internal/entities"
)

type FindByZipCodeUseCaseMock struct {
	mock.Mock
}

func (m *FindByZipCodeUseCaseMock) Execute(city string) (*entities.Location, error) {
	args := m.Called(city)
	return args.Get(0).(*entities.Location), args.Error(1)
}

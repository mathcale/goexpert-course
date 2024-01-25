package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/mathcale/goexpert-course/cloudrun/internal/entities"
)

type FindByCityNameUseCaseMock struct {
	mock.Mock
}

func (m *FindByCityNameUseCaseMock) Execute(city string) (*entities.Climate, error) {
	args := m.Called(city)
	return args.Get(0).(*entities.Climate), args.Error(1)
}

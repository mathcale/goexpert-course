package mocks

import (
	"github.com/mathcale/goexpert-course/cloudrun/internal/pkg/httpclient"
	"github.com/stretchr/testify/mock"
)

type HttpClientMock struct {
	mock.Mock
}

func (m *HttpClientMock) Get(endpoint string, responseObj interface{}) *httpclient.HttpClientError {
	args := m.Called(endpoint, responseObj)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*httpclient.HttpClientError)
}

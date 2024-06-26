package mocks

import (
	"context"

	"github.com/mathcale/goexpert-course/otel-lab/internal/pkg/httpclient"
	"github.com/stretchr/testify/mock"
)

type HttpClientMock struct {
	mock.Mock
}

func (m *HttpClientMock) Get(ctx context.Context, endpoint string, responseObj interface{}) *httpclient.HttpClientError {
	args := m.Called(ctx, endpoint, responseObj)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*httpclient.HttpClientError)
}

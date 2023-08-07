package mocks

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"github.com/stretchr/testify/mock"
)

type CourseRepositoryMock struct {
	mock.Mock
}

func (m *CourseRepositoryMock) Save(_ context.Context, c domain.Course) error {
	args := m.Called(c)
	return args.Error(0)
}

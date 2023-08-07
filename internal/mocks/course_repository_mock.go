package mocks

import (
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"github.com/stretchr/testify/mock"
)

type CourseRepositoryMock struct {
	mock.Mock
}

func (m *CourseRepositoryMock) Save(c domain.Course) error {
	args := m.Called(c)
	return args.Error(0)
}

package dummy

import (
	"context"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
)

type DummyCourseRepository struct{}

func (d *DummyCourseRepository) Save(ctx context.Context, course domain.Course) error {
	return nil
}

func NewDummyCourseRepository() *DummyCourseRepository {
	return &DummyCourseRepository{}
}

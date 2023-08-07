package domain

import "context"

type CourseRepository interface {
	Save(ctx context.Context, c Course) error
}

package domain

type CourseRepository interface {
	Save(c Course) error
}

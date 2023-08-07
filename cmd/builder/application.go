package builder

import (
	"github.com/dasalgadoc/clean-architecture-go/internal/application"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/storage/mysql"
)

type (
	Application struct {
		CourseCreator application.CourseCreator
	}

	repositories struct {
		CourseRepository domain.CourseRepository
	}
)

func BuildApplication() (*Application, error) {
	repos, err := buildRepositories()
	if err != nil {
		return nil, err
	}

	return &Application{
		CourseCreator: application.NewCourseCreator(repos.CourseRepository),
	}, nil
}

func buildRepositories() (*repositories, error) {
	user := "root"
	password := "root"
	host := "localhost"
	port := "3306"
	database := "clean_architecture_go"

	course, err := mysql.NewMysqlCourseRepository(user, password, host, port, database)
	if err != nil {
		return nil, err
	}

	return &repositories{
		CourseRepository: course,
	}, nil
}

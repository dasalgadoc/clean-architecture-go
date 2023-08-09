package builder

import (
	"github.com/dasalgadoc/clean-architecture-go/internal/application"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/bus/memory"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/storage/dummy"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/storage/mysql"
	"github.com/dasalgadoc/clean-architecture-go/shared/command"
	"github.com/dasalgadoc/clean-architecture-go/shared/event"
	"time"
)

type (
	Application struct {
		CourseCreator application.CourseCreator
		CommandBus    command.CommandBus
		EventBus      event.EventBus
	}

	repositories struct {
		CourseRepository domain.CourseRepository
	}
)

func BuildApplication() (*Application, error) {
	repos, err := buildDummyRepositories()
	if err != nil {
		return nil, err
	}

	eventBus := memory.NewEventBus()
	eventBus.Subscribe(domain.COURSE_CREATED_EVENT_NAME,
		application.NewCourseCounter())

	return &Application{
		CourseCreator: application.NewCourseCreator(repos.CourseRepository, eventBus),
		CommandBus:    memory.NewCommandBus(),
		EventBus:      eventBus,
	}, nil
}

func buildRepositories() (*repositories, error) {
	user := "root"
	password := "root"
	host := "localhost"
	port := "3306"
	database := "clean_architecture_go"
	timeout := 5 * time.Second

	course, err := mysql.NewMysqlCourseRepository(user, password, host, port, database, timeout)
	if err != nil {
		return nil, err
	}

	return &repositories{
		CourseRepository: course,
	}, nil
}

func buildDummyRepositories() (*repositories, error) {
	return &repositories{
		CourseRepository: dummy.NewDummyCourseRepository(),
	}, nil
}

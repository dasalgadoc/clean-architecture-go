package builder

import (
	"github.com/dasalgadoc/clean-architecture-go/internal/application"
	"github.com/dasalgadoc/clean-architecture-go/internal/domain"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/bus/memory"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/storage/dummy"
	"github.com/dasalgadoc/clean-architecture-go/internal/platform/storage/mysql"
	"github.com/dasalgadoc/clean-architecture-go/shared/command"
	"github.com/dasalgadoc/clean-architecture-go/shared/event"
	"github.com/kelseyhightower/envconfig"
)

type (
	Application struct {
		Config        config
		CourseCreator application.CourseCreator
		CommandBus    command.CommandBus
		EventBus      event.EventBus
	}

	repositories struct {
		CourseRepository domain.CourseRepository
	}
)

func BuildApplication() (*Application, error) {
	appConfig, err := getConfiguration()
	if err != nil {
		return nil, err
	}

	repos, err := buildRepositories(appConfig)
	if err != nil {
		return nil, err
	}

	eventBus := memory.NewEventBus()
	eventBus.Subscribe(domain.COURSE_CREATED_EVENT_NAME,
		application.NewCourseCounter())

	return &Application{
		Config:        *appConfig,
		CourseCreator: application.NewCourseCreator(repos.CourseRepository, eventBus),
		CommandBus:    memory.NewCommandBus(),
		EventBus:      eventBus,
	}, nil
}

func getConfiguration() (*config, error) {
	var cfg config
	err := envconfig.Process("CLEAN", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func buildRepositories(config *config) (*repositories, error) {
	course, err := mysql.NewMysqlCourseRepository(config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbHost,
		config.DbName,
		config.DbTimeout)
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

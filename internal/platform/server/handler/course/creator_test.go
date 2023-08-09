package course

import (
	"bytes"
	"github.com/dasalgadoc/clean-architecture-go/internal/application"
	"github.com/dasalgadoc/clean-architecture-go/internal/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type courseCreatorTestScenario struct {
	router           *gin.Engine
	writer           *httptest.ResponseRecorder
	test             *testing.T
	courseRepository *mocks.CourseRepositoryMock
	eventBusMock     *mocks.EventBusMock

	body []byte
}

func TestGivenAnInvalidRequestReturnBadRequest(t *testing.T) {
	test := startCourseCreatorTestScenario(t)
	test.setup()
	test.givenAName("")
	test.whenCreateCourseExecuted()
	test.thenResponseShouldBeBadRequest()
}

func TestGivenAValidRequestReturnCreated(t *testing.T) {
	test := startCourseCreatorTestScenario(t)
	test.setup()
	test.givenAName("OOP")
	test.andCourseRepositoryReturnNoError()
	test.andEventBusReturnNoError()
	test.whenCreateCourseExecuted()
	test.thenResponseShouldBeCreated()
}

func startCourseCreatorTestScenario(t *testing.T) *courseCreatorTestScenario {
	t.Parallel()

	return &courseCreatorTestScenario{
		test:             t,
		courseRepository: new(mocks.CourseRepositoryMock),
		eventBusMock:     new(mocks.EventBusMock),
	}
}

func (c *courseCreatorTestScenario) setup() {
	useCase := application.NewCourseCreator(c.courseRepository, c.eventBusMock)
	target := NewPostCourseCreator(useCase)

	gin.SetMode(gin.TestMode)
	c.router = gin.Default()
	c.router.POST("/course", target.Do)
}

func (c *courseCreatorTestScenario) givenAName(name string) {
	if name != "" {
		c.body = []byte(`{"name":"` + name + `"}`)
	}
}

func (c *courseCreatorTestScenario) andCourseRepositoryReturnNoError() {
	c.courseRepository.On("Save", mock.Anything).Return(nil)
}

func (c *courseCreatorTestScenario) andEventBusReturnNoError() {
	c.eventBusMock.On("Publish", mock.Anything).Return(nil)
}

func (c *courseCreatorTestScenario) whenCreateCourseExecuted() {
	req, err := http.NewRequest(http.MethodPost, "/course", bytes.NewBuffer(c.body))
	require.NoError(c.test, err)

	c.writer = httptest.NewRecorder()
	c.router.ServeHTTP(c.writer, req)
}

func (c *courseCreatorTestScenario) thenResponseShouldBeBadRequest() {
	assert.Equal(c.test, http.StatusBadRequest, c.writer.Code)
}

func (c *courseCreatorTestScenario) thenResponseShouldBeCreated() {
	assert.Equal(c.test, http.StatusCreated, c.writer.Code)
}
